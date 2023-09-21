package handler

import (
	"bufio"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"

	"github.com/shwatanap/connect-tic-tac-toe/src/connect/adapter"
	gamev1 "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/game/v1"
	gamev1connect "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/game/v1/gamev1connect"
	matchingv1 "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/matching/v1"
	matchingv1connect "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/matching/v1/matchingv1connect"
	"github.com/shwatanap/connect-tic-tac-toe/src/model"
)

var (
	server = flag.String("server", "http://localhost:50051", "server address")
)

type TicTacToe struct {
	sync.RWMutex
	Started  bool
	Finished bool
	Me       *model.Player
	Room     *model.Room
	Game     *model.Game
}

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{}
}

func Run(t *TicTacToe) int {
	if err := PreRun(t); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func PreRun(t *TicTacToe) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpClient := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	// マッチング問い合わせ
	err := Matching(ctx, matchingv1connect.NewMatchingServiceClient(httpClient, *server), t)
	if err != nil {
		return err
	}

	// マッチングできたので盤面生成
	t.Game = model.NewGame(t.Me.Symbol)

	// 双方向ストリーミングでゲーム処理
	return ExecPlay(ctx, gamev1connect.NewGameServiceClient(httpClient, *server), t)
}

func Matching(ctx context.Context, client matchingv1connect.MatchingServiceClient, t *TicTacToe) error {
	// マッチングリクエスト
	stream, err := client.JoinRoom(ctx, &connect.Request[matchingv1.JoinRoomRequest]{})
	if err != nil {
		return err
	}
	defer stream.Close()

	fmt.Println("Requested matching...")

	// ストリーミングでレスポンスを受け取る
	for {
		if !stream.Receive() {
			return stream.Err()
		}

		if stream.Msg().GetStatus() == matchingv1.JoinRoomResponse_STATUS_MATCHED {
			// マッチング成立
			t.Room = adapter.Room(stream.Msg().GetRoom())
			t.Me = adapter.Player(stream.Msg().GetMe())
			fmt.Printf("Matched room_id=%d\n", stream.Msg().GetRoom().GetId())
			return nil
		} else if stream.Msg().GetStatus() == matchingv1.JoinRoomResponse_STATUS_WAITTING {
			// 待機中
			fmt.Println("Waiting mathing...")
		}
	}
}

func ExecPlay(ctx context.Context, client gamev1connect.GameServiceClient, t *TicTacToe) error {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	// 双方向ストリーミングを開始する
	stream := client.Play(c)
	defer stream.CloseRequest()

	go func() {
		err := Send(c, stream, t)
		if err != nil {
			cancel()
		}
	}()

	err := Receive(c, stream, t)
	if err != nil {
		cancel()
		return err
	}

	return nil
}

func Play(t *TicTacToe) (bool, error) {
	fmt.Print("Input Your Move (ex. A-1):")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	// 入力された手を解析する
	text := stdin.Text()
	x, y, err := parseInput(text)
	if err != nil {
		return false, err
	}
	isGameOver, err := model.Move(x-1, y-1, t.Me.Symbol, t.Game)
	if err != nil {
		return false, err
	}

	return isGameOver, nil
}

// `A-2`の形式で入力された手を (x, y)=(1,2) の形式に変換する
func parseInput(txt string) (int32, int32, error) {
	ss := strings.Split(txt, "-")
	if len(ss) != 2 {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}

	xs := ss[0]
	xrs := []rune(strings.ToUpper(xs))
	x := int32(xrs[0]-rune('A')) + 1

	if x < 1 || 8 < x {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}

	ys := ss[1]
	y, err := strconv.ParseInt(ys, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}
	if y < 1 || 8 < y {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}

	return x, int32(y), nil
}

func Receive(ctx context.Context, stream *connect.BidiStreamForClient[gamev1.PlayRequest, gamev1.PlayResponse], t *TicTacToe) error {
	for {
		// サーバーからのストリーミングを受け取る
		res, err := stream.Receive()
		if err != nil {
			return err
		}

		t.Lock()
		switch res.GetEvent().(type) {
		case *gamev1.PlayResponse_Waiting:
			// 開始待機中
		case *gamev1.PlayResponse_Ready:
			// 開始
			t.Started = true
			model.Display(t.Me.Symbol, t.Game)
		case *gamev1.PlayResponse_Move:
			// 手を打たれた
			symbol := adapter.Symbol(res.GetMove().GetPlayer().GetSymbol())
			if symbol != t.Me.Symbol {
				move := res.GetMove().GetMove()
				// クライアント側のゲーム情報に反映させる
				model.Move(move.GetX(), move.GetY(), symbol, t.Game)
				fmt.Print("Input Your Move (ex. A-1):")
			}
		case *gamev1.PlayResponse_Finished:
			// ゲームが終了した
			t.Finished = true

			// 勝敗を表示する
			winner := adapter.Symbol(res.GetFinished().Winner)
			fmt.Println("")
			if winner == model.None {
				fmt.Println("Draw!")
			} else if winner == t.Me.Symbol {
				fmt.Println("You Win!")
			} else {
				fmt.Println("You Lose!")
			}

			// ループを終了する
			t.Unlock()
			return nil
		}
		t.Unlock()

		select {
		case <-ctx.Done():
			// キャンセルされたので終了する
			return nil
		default:
		}
	}
}

func Send(ctx context.Context, stream *connect.BidiStreamForClient[gamev1.PlayRequest, gamev1.PlayResponse], t *TicTacToe) error {
	for {
		t.RLock()

		if t.Finished {
			// recieve側で終了されたので、send側も終了する
			t.RUnlock()
			return nil
		} else if !t.Started {
			// 未開始なので、開始リクエストを送る
			err := stream.Send(&gamev1.PlayRequest{
				RoomId: t.Room.ID,
				Player: adapter.PBPlayer(t.Me),
				Action: &gamev1.PlayRequest_Start{
					Start: &gamev1.PlayRequest_StartAction{},
				},
			})
			t.RUnlock()
			if err != nil {
				return err
			}

			for {
				// 相手が開始するまで待機する
				t.RLock()
				if t.Started {
					// 開始をrecieveした
					t.RUnlock()
					fmt.Println("READY GO!")
					break
				}
				t.RUnlock()
				fmt.Println("Waiting until opponent player ready")
				time.Sleep(1 * time.Second)
			}
		} else {
			// 対戦中

			t.RUnlock()
			// 手の入力を待機する
			fmt.Print("Input Your Move (ex. A-1):")
			stdin := bufio.NewScanner(os.Stdin)
			stdin.Scan()

			// 入力された手を解析する
			text := stdin.Text()
			x, y, err := parseInput(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// 手を打つ
			t.Lock()
			_, err = model.Move(x, y, t.Me.Symbol, t.Game)
			t.Unlock()
			if err != nil {
				fmt.Println(err)
				continue
			}

			go func() {
				// サーバーに手を送る
				err = stream.Send(&gamev1.PlayRequest{
					RoomId: t.Room.ID,
					Player: adapter.PBPlayer(t.Me),
					Action: &gamev1.PlayRequest_Move{
						Move: &gamev1.PlayRequest_MoveAction{
							Move: &gamev1.Move{
								X: x,
								Y: y,
							},
						},
					},
				})
				if err != nil {
					fmt.Println(err)
				}
			}()

			// 一度手を打ったら5秒間待機する
			ch := make(chan int)
			go func(ch chan int) {
				fmt.Println("")
				for i := 0; i < 5; i++ {
					fmt.Printf("freezing in %d second.\n", (5 - i))
					time.Sleep(1 * time.Second)
				}
				fmt.Println("")
				ch <- 0
			}(ch)
			<-ch
		}

		select {
		case <-ctx.Done():
			// キャンセルされたので終了する
			return nil
		default:
		}
	}
}

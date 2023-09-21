package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/bufbuild/connect-go"

	"github.com/shwatanap/connect-tic-tac-toe/src/adapter"
	gamev1 "github.com/shwatanap/connect-tic-tac-toe/src/api/game/v1"
	"github.com/shwatanap/connect-tic-tac-toe/src/model"
)

type GameHandler struct {
	sync.RWMutex
	games  map[int32]*model.Game                                                    // ゲーム情報（盤面など）を格納する
	client map[int32][]*connect.BidiStream[gamev1.PlayRequest, gamev1.PlayResponse] // 状態変更時にクライアントにストリーミングを返すために格納する
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		games:  make(map[int32]*model.Game),
		client: make(map[int32][]*connect.BidiStream[gamev1.PlayRequest, gamev1.PlayResponse]),
	}
}

func (h *GameHandler) Play(ctx context.Context, stream *connect.BidiStream[gamev1.PlayRequest, gamev1.PlayResponse]) error {
	for {
		//クライアントからリクエストを受信したらreqにリクエストが代入されます
		req, err := stream.Receive()
		if err != nil {
			return err
		}

		roomID := req.GetRoomId()
		player := adapter.Player(req.GetPlayer())

		//oneofで複数の型のリクエストがくるのでswtich文で処理します
		switch req.GetAction().(type) {
		case *gamev1.PlayRequest_Start:
			//ゲーム開始リクエスト
			err := h.start(stream, roomID, player)
			if err != nil {
				return err
			}
		case *gamev1.PlayRequest_Move:
			//石を置いた時のリクエスト
			action := req.GetMove()
			x := action.GetMove().GetX()
			y := action.GetMove().GetY()
			err := h.move(roomID, x, y, player)
			if err != nil {
				return err
			}
		}
	}
}

func (h *GameHandler) start(stream *connect.BidiStream[gamev1.PlayRequest, gamev1.PlayResponse], roomID int32, me *model.Player) error {
	h.Lock()
	defer h.Unlock()

	//ゲーム情報がなければ作成する
	g := h.games[roomID]
	if g == nil {
		g = model.NewGame(model.None)
		h.games[roomID] = g
		h.client[roomID] = make([]*connect.BidiStream[gamev1.PlayRequest, gamev1.PlayResponse], 0, 2)
	}

	//自分のクライアントを格納する
	h.client[roomID] = append(h.client[roomID], stream)

	if len(h.client[roomID]) == 2 {
		// 二人揃ったので開始する
		for _, s := range h.client[roomID] {
			// クライアントにゲーム開始を通知する
			err := s.Send(&gamev1.PlayResponse{
				Event: &gamev1.PlayResponse_Ready{
					Ready: &gamev1.PlayResponse_ReadyEvent{},
				},
			})
			if err != nil {
				return err
			}
		}
		fmt.Printf("game has started room_id=%v\n", roomID)
	} else {
		//まだ揃ってないので待機中であることをクライアントに通知する
		err := stream.Send(&gamev1.PlayResponse{
			Event: &gamev1.PlayResponse_Waiting{
				Waiting: &gamev1.PlayResponse_WaitingEvent{},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *GameHandler) move(roomID int32, x int32, y int32, p *model.Player) error {
	h.Lock()
	defer h.Unlock()

	g := h.games[roomID]

	finished, err := model.Move(x, y, p.Symbol, g)
	if err != nil {
		return err
	}

	for _, s := range h.client[roomID] {
		// 手が打たれたことをクライアントに通知する
		err := s.Send(&gamev1.PlayResponse{
			Event: &gamev1.PlayResponse_Move{
				Move: &gamev1.PlayResponse_MoveEvent{
					Player: adapter.PBPlayer(p),
					Move: &gamev1.Move{
						X: x,
						Y: y,
					},
					Board: adapter.PBBoard(g.Board),
				},
			},
		})
		if err != nil {
			return err
		}

		if finished {
			// ゲーム終了通知する
			err := s.Send(
				&gamev1.PlayResponse{
					Event: &gamev1.PlayResponse_Finished{
						Finished: &gamev1.PlayResponse_FinishedEvent{
							Winner: adapter.PBSymbol(model.WhoWinner(g)),
							Board:  adapter.PBBoard(g.Board),
						},
					},
				},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

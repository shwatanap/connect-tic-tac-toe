package handler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shwatanap/connect-tic-tac-toe/src/model"
	"github.com/shwatanap/connect-tic-tac-toe/src/standard/adapter"
	matchingv1 "github.com/shwatanap/connect-tic-tac-toe/src/standard/api/matching/v1"
)

type MatchingHandler struct {
	sync.RWMutex
	Rooms       map[int32]*model.Room
	maxPlayerID int32
	matchingv1.UnimplementedMatchingServiceServer
}

func NewMatchingHandler() *MatchingHandler {
	return &MatchingHandler{
		Rooms: make(map[int32]*model.Room),
	}
}

func (h *MatchingHandler) JoinRoom(req *matchingv1.JoinRoomRequest, stream matchingv1.MatchingService_JoinRoomServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), 2*time.Minute)
	defer cancel()

	h.Lock()

	// プレイヤーの新規作成
	h.maxPlayerID++
	me := &model.Player{
		ID: h.maxPlayerID,
	}

	// 空いている部屋を探す
	for _, room := range h.Rooms {
		if room.Guest == nil {
			me.Symbol = model.Cross
			room.Guest = me
			stream.Send(&matchingv1.JoinRoomResponse{
				Status: matchingv1.JoinRoomResponse_STATUS_MATCHED,
				Room:   adapter.PBRoom(room),
				Me:     adapter.PBPlayer(room.Guest),
			})
			h.Unlock()
			fmt.Printf("matched room_id=%v\n", room.ID)
			return nil
		}
	}

	// 空いている部屋がなかったので部屋を作る
	me.Symbol = model.Circle
	room := &model.Room{
		ID:   int32(len(h.Rooms)) + 1,
		Host: me,
	}
	h.Rooms[room.ID] = room
	h.Unlock()

	stream.Send(&matchingv1.JoinRoomResponse{
		Status: matchingv1.JoinRoomResponse_STATUS_WAITTING,
		Room:   adapter.PBRoom(room),
	})

	ch := make(chan int)
	go func(ch chan<- int) {
		for {
			h.RLock()
			guest := room.Guest
			h.RUnlock()

			if guest != nil {
				stream.Send(&matchingv1.JoinRoomResponse{
					Status: matchingv1.JoinRoomResponse_STATUS_MATCHED,
					Room:   adapter.PBRoom(room),
					Me:     adapter.PBPlayer(room.Host),
				})
				ch <- 0
				break
			}
			time.Sleep(1 * time.Second)

			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		return status.Errorf(codes.DeadlineExceeded, "マッチングできませんでした")
	}

	return nil
}

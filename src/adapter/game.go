package adapter

import (
	"fmt"

	gamev1 "github.com/shwatanap/connect-tic-tac-toe/src/api/game/v1"
	mactingv1 "github.com/shwatanap/connect-tic-tac-toe/src/api/matching/v1"
	"github.com/shwatanap/connect-tic-tac-toe/src/model"
)

func Room(r *mactingv1.Room) *model.Room {
	return &model.Room{
		ID:    r.GetId(),
		Host:  Player(r.GetHost()),
		Guest: Player(r.GetGuest()),
	}
}

func Player(p *gamev1.Player) *model.Player {
	return &model.Player{
		ID:     p.GetId(),
		Symbol: Symbol(p.GetSymbol()),
	}
}

func Symbol(c gamev1.Symbol) model.Symbol {
	switch c {
	case gamev1.Symbol_EMPTY:
		return model.Empty
	case gamev1.Symbol_CIRCLE:
		return model.Circle
	case gamev1.Symbol_CROSS:
		return model.Cross
	case gamev1.Symbol_NONE:
		return model.None
	}

	panic(fmt.Sprintf("unknwon symbol=%v", c))
}

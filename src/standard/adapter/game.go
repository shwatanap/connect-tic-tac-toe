package adapter

import (
	"fmt"

	"github.com/shwatanap/connect-tic-tac-toe/src/model"
	gamev1 "github.com/shwatanap/connect-tic-tac-toe/src/standard/api/game/v1"
	mactingv1 "github.com/shwatanap/connect-tic-tac-toe/src/standard/api/matching/v1"
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
	case gamev1.Symbol_SYMBOL_EMPTY:
		return model.Empty
	case gamev1.Symbol_SYMBOL_CIRCLE:
		return model.Circle
	case gamev1.Symbol_SYMBOL_CROSS:
		return model.Cross
	case gamev1.Symbol_SYMBOL_NONE:
		return model.None
	}

	panic(fmt.Sprintf("unknwon symbol=%v", c))
}

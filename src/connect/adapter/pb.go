package adapter

import (
	gamev1 "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/game/v1"
	mactingv1 "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/matching/v1"
	"github.com/shwatanap/connect-tic-tac-toe/src/model"
)

func PBRoom(r *model.Room) *mactingv1.Room {
	return &mactingv1.Room{
		Id:    r.ID,
		Host:  PBPlayer(r.Host),
		Guest: PBPlayer(r.Guest),
	}
}

func PBPlayer(p *model.Player) *gamev1.Player {
	if p == nil {
		return nil
	}
	return &gamev1.Player{
		Id:     p.ID,
		Symbol: PBSymbol(p.Symbol),
	}
}

func PBSymbol(c model.Symbol) gamev1.Symbol {
	switch c {
	case model.Circle:
		return gamev1.Symbol_CIRCLE
	case model.Cross:
		return gamev1.Symbol_CROSS
	case model.Empty:
		return gamev1.Symbol_EMPTY
	}

	return gamev1.Symbol_UNKNOWN
}

func PBBoard(b *model.Board) *gamev1.Board {
	pbCols := make([]*gamev1.Board_Sym, 0, 10)

	for _, col := range b.Cells {
		pbCells := make([]gamev1.Symbol, 0, 10)
		for _, c := range col {
			pbCells = append(pbCells, PBSymbol(c))
		}
		pbCols = append(pbCols, &gamev1.Board_Sym{
			Cells: pbCells,
		})
	}

	return &gamev1.Board{
		Cols: pbCols,
	}
}

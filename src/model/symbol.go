package model

type Symbol int

const (
	Empty  Symbol = iota // 誰も打ってない
	Circle               // マル
	Cross                // バツ
	None                 // 引き分け判定用
)

// マーカーを文字列に変換します
func SymbolToStr(se Symbol) string {
	switch se {
	case Circle:
		return "○"
	case Cross:
		return "×"
	case Empty:
		return " "
	}

	return ""
}

// 対戦相手のマーカーを取得します
func OpponentSymbol(me Symbol) Symbol {
	switch me {
	case Circle:
		return Cross
	case Cross:
		return Circle
	}

	panic("invalid state")
}

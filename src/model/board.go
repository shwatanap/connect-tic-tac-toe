package model

import "fmt"

type Board struct {
	Cells [][]Symbol
}

// 勝ち手
var Lines = [][][]int{
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

func NewBoard() *Board {
	// 3x3のマスの盤面を二次元配列で作成
	b := &Board{
		Cells: make([][]Symbol, 3),
	}
	for i := 0; i < 3; i++ {
		b.Cells[i] = make([]Symbol, 3)
	}

	return b
}

// マーカーを書く
func PutMarker(x int32, y int32, s Symbol, b *Board) error {
	// そのマスにマーカーを書けるかチェック
	if !CanPutMarker(x, y, b) {
		return fmt.Errorf("can not put Marker x=%v, y=%v symbol=%v", x, y, SymbolToStr(s))
	}

	// マスにマーカーを書く
	b.Cells[x][y] = s

	return nil
}

// マスに石を置けるか判定します
func CanPutMarker(x int32, y int32, b *Board) bool {
	// すでに石がある場合は石を置けません
	return b.Cells[x][y] == Empty
}

// 空きマスがあるか判定
func IsAvailableEmpty(b *Board) bool {
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			// 空きマスが一つでもあればtrue
			if b.Cells[i][j] == Empty {
				return true
			}
		}
	}

	// 空きマスが一つもなければfalse
	return false
}

// ラインができているか判定
func IsAvailableLine(s Symbol, b *Board) bool {
	for _, line := range Lines {
		for i, cell := range line {
			// 勝ち手に合わない場合にbreak
			if b.Cells[cell[0]][cell[1]] != s {
				break
			}

			// 最後まで勝ち手に合えばtrueを返す
			if i == len(line)-1 {
				return true
			}
		}
	}

	// 一つも勝ち手に当てはまらなければfalseを返す
	return false
}

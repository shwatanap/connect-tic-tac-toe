package model

import "fmt"

type Game struct {
	Board    *Board
	Finished bool
	Me       Symbol
}

func NewGame(me Symbol) *Game {
	return &Game{
		Board: NewBoard(),
		Me:    me,
	}
}

// 手を打ち、その後盤面を出力します。
// 返り値として、ゲームが終了したかを返す。
func Move(x int32, y int32, s Symbol, g *Game) (bool, error) {
	if g.Finished {
		return true, nil
	}
	err := PutMarker(x-1, y-1, s, g.Board)
	if err != nil {
		return false, err
	}
	Display(g.Me, g)
	if IsGameOver(g) != NoWin {
		fmt.Println("finished")
		g.Finished = true
		return true, nil
	}

	return false, nil
}

// ゲームが終了したか判定
// マルとバツ双方に置ける場所がなければ終了
func IsGameOver(g *Game) Winner {
	if !IsAvailableEmpty(g.Board) {
		return Draw
	}

	if IsAvailableLine(Circle, g.Board) {
		return CircleWin
	}

	if IsAvailableLine(Cross, g.Board) {
		return CrossWin
	}

	return NoWin
}

// 勝者のマーカーを返す
// 引き分けの場合はNoneを返す
func WhoWinner(g *Game) Symbol {
	if IsAvailableLine(Circle, g.Board) {
		return Circle
	}

	if IsAvailableLine(Cross, g.Board) {
		return Cross
	}

	return None
}

// 盤面の出力
func Display(me Symbol, g *Game) {
	fmt.Println("")
	if me != None {
		fmt.Printf("You: %v\n", SymbolToStr(me))
	}

	fmt.Print(" ｜")
	rs := []rune("ABC")
	for i, r := range rs {
		fmt.Printf("%v", string(r))
		if i < len(rs)-1 {
			fmt.Print("｜")
		}
	}
	fmt.Print("｜")
	fmt.Print("\n")
	fmt.Println("ーーーーーー")

	for j := 0; j < 3; j++ {
		fmt.Printf("%d", j+1)
		fmt.Print("｜")
		for i := 0; i < 3; i++ {
			fmt.Print(SymbolToStr(g.Board.Cells[i][j]))
			fmt.Print("｜")
		}
		fmt.Print("\n")
	}

	fmt.Println("ーーーーーー")

	fmt.Print("\n")
}

package main

import (
	"os"

	handler "github.com/shwatanap/connect-tic-tac-toe/src/connect/handler/client"
)

func main() {
	os.Exit(handler.Run(handler.NewTicTacToe()))
}

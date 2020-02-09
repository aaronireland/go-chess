package main

import (
	"fmt"

	"github.com/aaronireland/go-chess/pkg/chess"
)

func main() {
	board, err := chess.NewBoard()

	if err != nil {
		panic(err)
	}

	fmt.Println(board)
	fmt.Println("White to move... e4")
	board.MovePieceAlgebraic(chess.WhitePawn.Index, "e2", "e4")
	fmt.Println(board)
}

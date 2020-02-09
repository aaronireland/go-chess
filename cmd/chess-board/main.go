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
	board.MovePieceAlgebraic(board.GetIndex(chess.Pawn(chess.WHITE)), "e2", "e4")
	fmt.Println(board)
}

package chess

import (
	"fmt"
	"strings"
	"testing"
)

type test struct {
	Condition   bool
	ShouldPass  bool
	Description string
	Err         error
}
type tests []test

func (tests tests) Run(t *testing.T) {
	for _, test := range tests {
		if !(test.Condition == test.ShouldPass) {
			if test.Err != nil {
				t.Errorf("FAILED: %s: %s", test.Description, test.Err)
			} else {
				t.Errorf("FAILED: %s", test.Description)
			}
		}
	}
}

var boardStr = []string{
	" A B C D E F G H",
	"8♜ ♞ ♝ ♛ ♚ ♝ ♞ ♜",
	"7♟ ♟ ♟ ♟ ♟ ♟ ♟ ♟",
	"6- - - - - - - -",
	"5- - - - - - - -",
	"4- - - - - - - -",
	"3- - - - - - - -",
	"2♙ ♙ ♙ ♙ ♙ ♙ ♙ ♙",
	"1♖ ♘ ♗ ♕ ♔ ♗ ♘ ♖",
	"",
}

func TestNewBoard(t *testing.T) {
	board, err := NewBoard()

	boardTests := tests{
		test{err == nil, true, "Making a new board should not result in error", err},
		test{len(board.Positions) == len(board.Pieces), true, "Number of piece types should match number of bitboards", nil},
		test{board.Occupied.Population() == 32, true, "A new board should have 32 pieces", nil},
	}

	board, err = NewBoard(emptyBoard, emptyBoard)
	boardTests = append(boardTests, test{err != nil, true, "Attempting to make a board copy without enough bitboards should error", err})

	var bitboards []Bitboard
	for i := 0; i < (FILES + 1); i++ {
		bitboards = append(bitboards, emptyBoard)
	}
	board, err = NewBoard(bitboards...)
	boardTests = append(boardTests, test{err != nil, true, "Attempting to make a board copy with too many bitboards should error", err})

	bitboards = []Bitboard{}
	for i := 0; i < (len(Pieces)); i++ {
		bitboards = append(bitboards, emptyBoard)
	}

	board, err = NewBoard(bitboards...)
	boardTests = append(boardTests, test{err == nil, true, "Attempting to make a board copy should not error", err})
	boardTests = append(boardTests, test{board.Occupied.Population() == 0, true, "Empty board should have no pieces", err})

	boardTests.Run(t)

}

func TestGetSquare(t *testing.T) {
	bitboards := []Bitboard{
		initWhiteRooks, emptyBoard, emptyBoard, emptyBoard, emptyBoard, emptyBoard,
		emptyBoard, emptyBoard, emptyBoard, emptyBoard, emptyBoard, emptyBoard,
	}

	board, err := NewBoard(bitboards...)
	if err != nil {
		t.Errorf("Unexpected error generating new board with only white rooks: %s", err)
	}
	a1occupied, a1piece := board.GetSquare(0)
	a8occupied, a8piece := board.GetSquare(7)
	squareTests := tests{
		test{board.Occupied.Population() == 2, true, "Board with white rooks should have 2 pieces", err},
		test{a1occupied, true, "Square a1 should be occupied", nil},
		test{a8occupied, true, "Square a8 should be occupied", nil},
		test{a1piece.Equals(WhiteRook), true, "The piece at a1 should be a white rook", nil},
		test{a8piece.Equals(WhiteRook), true, "The piece at a8 should be a white rook", nil},
	}

	for i := 0; i < (RANKS * FILES); i++ {
		if i == 0 || i == 7 {
			continue
		}
		occupied, _ := board.GetSquare(i)
		squareTests = append(squareTests, test{occupied, false, fmt.Sprintf("Square at index %d should be unoccupied", i), nil})
	}

	squareTests.Run(t)

}

func TestString(t *testing.T) {
	board, _ := NewBoard()
	actual := strings.Split(board.String(), "\n")

	if len(actual) < len(boardStr) {
		t.Errorf("board.String() missing lines, %d expected, %d actual", len(boardStr), len(actual))
	} else if len(actual) > len(boardStr) {
		t.Errorf("expected string missing lines, %d expected, %d actual", len(boardStr), len(actual))
	}
	for i, line := range actual {
		if line != actual[i] {
			t.Errorf("board.String() expected %s, actual: %s", line, actual[i])
		}
	}
}

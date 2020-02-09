package chess

import (
	"fmt"
	"sort"
)

// Color is an enum to distinguish white/black pieces
type Color uint8

// Symbol is a enum for notation for the chess pieces
type Symbol rune

// WHITE is 0
// BLACK is 1
// PAWN is an empty rune
// BISHOP is B
// KNIGHT is N
const (
	WHITE Color = 0
	BLACK Color = 1

	PAWN   Symbol = 0
	BISHOP Symbol = 'B'
	KNIGHT Symbol = 'N'
	ROOK   Symbol = 'R'
	QUEEN  Symbol = 'Q'
	KING   Symbol = 'K'
)

// PieceNames maps the symbol to the descriptive name
var PieceNames = map[Symbol]string{
	PAWN:   "pawn",
	BISHOP: "bishop",
	KNIGHT: "knight",
	ROOK:   "rook",
	QUEEN:  "queen",
	KING:   "king",
}

// WhitePawn WhiteBishop etc...  are Piece objects with attribues unique to
// the particular piece type
var (
	WhitePawn   = pawn(WHITE, 0)
	WhiteBishop = bishop(WHITE, 1)
	WhiteKnight = knight(WHITE, 2)
	WhiteRook   = rook(WHITE, 3)
	WhiteQueen  = queen(WHITE, 4)
	WhiteKing   = king(WHITE, 5)
	BlackPawn   = pawn(BLACK, 6)
	BlackBishop = bishop(BLACK, 7)
	BlackKnight = knight(BLACK, 8)
	BlackRook   = rook(BLACK, 9)
	BlackQueen  = queen(BLACK, 10)
	BlackKing   = king(BLACK, 11)
)

// Pieces is a slice of all chess piece types. A slice is used instead of a map
// to avoid overhead in lookups
var Pieces = []Piece{
	WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhitePawn,
	BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackPawn,
}

func init() {
	// Ensure the Pieces are sorted by index
	sort.SliceStable(Pieces, func(i, j int) bool {
		return Pieces[i].Index < Pieces[j].Index
	})

}

// PieceUnicodes provides a visual representation for each piece for drawing boards/icons in a console
var PieceUnicodes = map[Piece]string{
	WhitePawn:   "♙",
	WhiteBishop: "♗",
	WhiteKnight: "♘",
	WhiteRook:   "♖",
	WhiteQueen:  "♕",
	WhiteKing:   "♔",
	BlackPawn:   "♟",
	BlackBishop: "♝",
	BlackKnight: "♞",
	BlackRook:   "♜",
	BlackQueen:  "♛",
	BlackKing:   "♚",
}

// Piece groups the attributes required to distinguish pieces
type Piece struct {
	Color  Color
	Symbol Symbol
	Value  uint8
	Index  int
}

func (c Color) String() string {
	if c == WHITE {
		return "white"
	} else if c == BLACK {
		return "black"
	}

	return ""
}

func (s Symbol) String() string {
	if name, ok := PieceNames[s]; ok {
		return name
	}

	return ""
}

func (p Piece) String() string {
	return fmt.Sprintf("%s %s", p.Color, p.Symbol)
}

// Unicode returns the unicode symbol/icon for the piece
func (p Piece) Unicode() string {
	for piece, unicode := range PieceUnicodes {
		if piece.Equals(p) {
			return unicode
		}
	}

	return ""
}

// SameType checks piece equality regardless of color
func (p Piece) SameType(p2 Piece) bool {
	return p.Symbol == p2.Symbol
}

// Equals checks piece equality (color and type)
func (p Piece) Equals(p2 Piece) bool {
	return p.Color == p2.Color && p.SameType(p2)
}

func pawn(c Color, i int) Piece {
	return Piece{c, PAWN, 1, i}
}

func bishop(c Color, i int) Piece {
	return Piece{c, BISHOP, 3, i}
}

func knight(c Color, i int) Piece {
	return Piece{c, KNIGHT, 3, i}
}

func rook(c Color, i int) Piece {
	return Piece{c, ROOK, 5, i}
}

func queen(c Color, i int) Piece {
	return Piece{c, QUEEN, 9, i}
}

func king(c Color, i int) Piece {
	return Piece{c, KING, 0, i}
}

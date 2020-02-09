package chess

import "fmt"

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

// PieceUnicodes provides a visual representation for each piece for drawing boards/icons in a console
var PieceUnicodes = map[Piece]string{
	Pawn(WHITE):   "♙",
	Bishop(WHITE): "♗",
	Knight(WHITE): "♘",
	Rook(WHITE):   "♖",
	Queen(WHITE):  "♕",
	King(WHITE):   "♔",
	Pawn(BLACK):   "♟",
	Bishop(BLACK): "♝",
	Knight(BLACK): "♞",
	Rook(BLACK):   "♜",
	Queen(BLACK):  "♛",
	King(BLACK):   "♚",
}

// Piece groups the attributes required to distinguish pieces
type Piece struct {
	Color  Color
	Symbol Symbol
	Value  uint8
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

// Pawn returns a piece representing a pawn of the requested color
func Pawn(c Color) Piece {
	return Piece{c, PAWN, 1}
}

// Bishop returns a piece representing a bishop of the requested color
func Bishop(c Color) Piece {
	return Piece{c, BISHOP, 3}
}

// Knight returns a piece representing a knight of the requested color
func Knight(c Color) Piece {
	return Piece{c, KNIGHT, 3}
}

// Rook returns a piece representing a rook of the requested color
func Rook(c Color) Piece {
	return Piece{c, ROOK, 5}
}

// Queen returns a piece representing a queen of the requested color
func Queen(c Color) Piece {
	return Piece{c, QUEEN, 9}
}

// King returns a piece representing a king of the requested color
func King(c Color) Piece {
	return Piece{c, KING, 0}
}

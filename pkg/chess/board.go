package chess

import "fmt"

// RANKS defines the number of rows on the board
// FILES defines the number of columns on the board
const (
	RANKS uint8 = 8
	FILES uint8 = 8
)

const (
	initWhiteRooks   = Bitboard(0x0000000000000081)
	initWhiteKnights = Bitboard(0x0000000000000042)
	initWhiteBishops = Bitboard(0x0000000000000024)
	initWhiteQueen   = Bitboard(0x0000000000000008)
	initWhiteKing    = Bitboard(0x0000000000000010)
	initWhitePawns   = Bitboard(0x000000000000ff00)
	initBlackRooks   = Bitboard(0x8100000000000000)
	initBlackKnights = Bitboard(0x4200000000000000)
	initBlackBishops = Bitboard(0x2400000000000000)
	initBlackQueen   = Bitboard(0x0800000000000000)
	initBlackKing    = Bitboard(0x1000000000000000)
	initBlackPawns   = Bitboard(0x00ff000000000000)
)

// Board is a type of piece-centric representation of a chess board referred to a Bitboard.
// For more information see: https://www.chessprogramming.org/Bitboards
type Board struct {
	Positions []Bitboard
	Pieces    []Piece
	Occupied  Bitboard // Union of all piece positions gives current occupied squares
}

// NewBoard returns a new instance of the chess board optionally initialized with the requested board positions
func NewBoard(positions ...Bitboard) (*Board, error) {
	board := Board{
		Pieces: []Piece{
			Rook(WHITE), Knight(WHITE), Bishop(WHITE), Queen(WHITE), King(WHITE), Pawn(WHITE),
			Rook(BLACK), Knight(BLACK), Bishop(BLACK), Queen(BLACK), King(BLACK), Pawn(BLACK),
		},
	}
	if len(positions) > 0 && len(positions) != len(board.Pieces) {
		err := fmt.Errorf(
			"Unable to determine board position, expecting %d bitboards, received %d",
			len(board.Positions), len(positions),
		)
		return nil, err
	} else if len(positions) > 0 {
		board.Positions = make([]Bitboard, len(positions))
		copy(board.Positions, positions)
	} else {
		board.Positions = []Bitboard{
			initWhiteRooks, initWhiteKnights, initWhiteBishops, initWhiteQueen, initWhiteKing, initWhitePawns,
			initBlackRooks, initBlackKnights, initBlackBishops, initBlackQueen, initBlackKing, initBlackPawns,
		}
	}

	board.Occupied = Union(board.Positions...)

	return &board, nil
}

// GetSquare gives whether a square is occupied and if so by which piece for a given index 0-63
func (b Board) GetSquare(index int) (bool, *Piece) {
	if b.Occupied.GetBit(index) != 0 { // If 0, this square is unoccupied
		for i := 0; i < len(b.Positions); i++ {
			if b.Positions[i].GetBit(index) != 0 {
				return true, &b.Pieces[i]
			}
		}
	}
	return false, nil

}

// GetIndex returns the index of the requested Piece or -1 if not found
func (b Board) GetIndex(p Piece) int {
	for index, piece := range b.Pieces {
		if piece.Equals(p) {
			return index
		}
	}
	return -1
}

// String displays the current board a simple console-friendly unicode grid
func (b *Board) String() string {
	s := "\n A B C D E F G H\n"
	for rank := int(RANKS - 1); rank >= 0; rank-- {
		s += fmt.Sprintf("%d", rank+1)
		for file := 0; file < int(FILES); file++ {
			index := (rank * int(FILES)) + file
			if occupied, piece := b.GetSquare(index); occupied {
				s += piece.Unicode()
			} else {
				s += "-"
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

/******************************************************************************
*                   Place, Remove, and Move Pieces
******************************************************************************/

// PlacePiece marks the provided position index (0-63) as occupied on the bitboard
// defined by the piece index (i.e. Board.Position[int]) for this board
func (b *Board) PlacePiece(piece, position int) {
	b.Positions[piece].SetBit(position)
	b.Occupied.SetBit(position)
}

// RemovePiece marks the provided position index (0-63) as unoccupied  on the
// bitboard defined by the piece index (i.e. Board.Position[int]) for this board
func (b *Board) RemovePiece(piece, position int) {
	b.Positions[piece].ClearBit(position)
	b.Occupied.ClearBit(position)
}

// MovePiece updates the bitboard specified at the given piece index to simulate
// a piece having moved (i.e. clears the bit at the first given index and sets the bit
// at the second given index)
func (b *Board) MovePiece(piece, from, to int) {
	b.RemovePiece(piece, from)
	b.PlacePiece(piece, to)
}

// PlacePieceAlgebraic updates the bitboard specified at the given piece index to
// simulate placing a piece using the letters a through h to mark files and 1
// through 8 to mark ranks
func (b *Board) PlacePieceAlgebraic(piece int, position string) {
	b.PlacePiece(piece, AlgebraicToBit(position))
}

// RemovePieceAlgebraic updates the bitboard specified at the given piece index to
// simulate removing a piece from the positions given using the letters a
// through h to mark files and 1 through 8 to mark ranks
func (b *Board) RemovePieceAlgebraic(piece int, position string) {
	b.RemovePiece(piece, AlgebraicToBit(position))
}

// MovePieceAlgebraic updates the bitboard specified at the given piece index to
// simulate moving a piece from and to the speicified positions given using the letters a
// through h to mark files and 1 through 8 to mark ranks
func (b *Board) MovePieceAlgebraic(piece int, from, to string) {
	b.RemovePieceAlgebraic(piece, from)
	b.PlacePieceAlgebraic(piece, to)
}

// PlacePieceCartesian updates the bitboard specified at the given piece index to
// simulate placing a piece at the position given using x and y coordinates 0 through 7
func (b *Board) PlacePieceCartesian(piece, x, y int) {
	b.PlacePiece(piece, CartesianToBit(x, y))
}

// RemovePieceCartesian updates the bitboard specified at the given piece index to
// simulate removing a piece at the position given using x and y coordinates 0 through 7
func (b *Board) RemovePieceCartesian(piece, x, y int) {
	b.RemovePiece(piece, CartesianToBit(x, y))
}

// MovePieceCartesian updates the bitboard at the given piece index to
// simulate moving a piece from and to the positions given using x and y coordnidates 0 through 7
func (b *Board) MovePieceCartesian(piece, fromX, fromY, toX, toY int) {
	b.RemovePiece(piece, CartesianToBit(fromX, fromY))
	b.PlacePiece(piece, CartesianToBit(toX, toY))
}

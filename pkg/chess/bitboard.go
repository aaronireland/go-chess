package chess

// Bitboard is a single 64-bit word/register used to represent the game state of a chess board
// using a little-endian mapping of bits to the rank/file coordinates of the board.
// For an 8x8 board, this mapping looks like this:
//
//  8 | 56 57 58 59 60 61 62 63
//  7 | 48 49 50 51 52 53 54 55
//  6 | 40 41 42 43 44 45 46 47
//  5 | 32 33 34 35 36 37 38 39
//  4 | 24 25 26 27 28 29 30 31
//  3 | 16 17 18 19 20 21 22 23
//  2 | 8  9  10 11 12 13 14 15
//  1 | 0  1  2  3  4  5  6  7
//    -------------------------
//      a  b  c  d  e  f  g  h
type Bitboard uint64

// SetBit sets to 1 the bit at the requested index
func (b *Bitboard) SetBit(index int) {
	var mask Bitboard = (1 << uint(index))
	*b |= mask
}

// ClearBit set to 0 the bit at the requested index
func (b *Bitboard) ClearBit(index int) {
	var mask Bitboard = ^(1 << uint(index))
	*b &= mask
}

// ToggleBit switches the bit at the requested index (1 -> 0 -> 1, etc)
func (b *Bitboard) ToggleBit(index int) {
	var mask Bitboard = (1 << uint(index))
	*b ^= mask
}

// GetBit returns the value of the bit at the requested index
func (b Bitboard) GetBit(index int) int {
	return int((b >> uint(index)) & 1)
}

// IsBitSet checks if a bit is set at the requested index
func (b Bitboard) IsBitSet(index int) bool {
	var mask Bitboard = 1 << Bitboard(index)
	return (b & mask) != 0
}

// Union overlays a slice of bitmaps into a single structure
func Union(bitmaps ...Bitboard) Bitboard {
	var allmaps Bitboard
	for _, b := range bitmaps {
		allmaps = allmaps | b
	}
	return allmaps
}

// Population calculates the population count (Hamming weight) of an integer
// using a divide-and-conquer approach.
//
// See <http://en.wikipedia.org/wiki/Hamming_weight> for a complete description
// of this implementation.
func (b Bitboard) Population() int {
	var mask1, mask2, mask4 Bitboard
	mask1 = 0x5555555555555555 // 0101...
	mask2 = 0x3333333333333333 // 00110011..
	mask4 = 0x0f0f0f0f0f0f0f0f // 00001111...
	b -= (b >> 1) & mask1
	b = (b & mask2) + ((b >> 2) & mask2)
	b = (b + (b >> 4)) & mask4
	b += b >> 8
	b += b >> 16
	b += b >> 32
	return int(b & 0x7f)
}

//-----------------------------------------------------------------------------
// Flipping and rotating
//-----------------------------------------------------------------------------

// These are Go ports of the functions given on the Chess Programming wiki:
// <https://chessprogramming.wikispaces.com/Flipping+Mirroring+and+Rotating>.

// FlipVertical returns a new bitboard flipped vertically about the center ranks.
func (b Bitboard) FlipVertical() Bitboard {
	k1 := Bitboard(0x00FF00FF00FF00FF)
	k2 := Bitboard(0x0000FFFF0000FFFF)
	b = ((b >> 8) & k1) | ((b & k1) << 8)
	b = ((b >> 16) & k2) | ((b & k2) << 16)
	b = (b >> 32) | (b << 32)
	return b
}

// FlipHorizontal returns a new bitboard flipped horizontally about the center files
func (b Bitboard) FlipHorizontal() Bitboard {
	k1 := Bitboard(0x5555555555555555)
	k2 := Bitboard(0x3333333333333333)
	k4 := Bitboard(0x0f0f0f0f0f0f0f0f)
	b = ((b >> 1) & k1) + 2*(b&k1)
	b = ((b >> 2) & k2) + 4*(b&k2)
	b = ((b >> 4) & k4) + 16*(b&k4)
	return b
}

// FlipDiagonalA1H8 returns a new bitboard flipped about the diagonal a1-h8.
func (b Bitboard) FlipDiagonalA1H8() Bitboard {
	var t Bitboard
	k1 := Bitboard(0x5500550055005500)
	k2 := Bitboard(0x3333000033330000)
	k4 := Bitboard(0x0f0f0f0f00000000)
	t = k4 & (b ^ (b << 28))
	b ^= t ^ (t >> 28)
	t = k2 & (b ^ (b << 14))
	b ^= t ^ (t >> 14)
	t = k1 & (b ^ (b << 7))
	b ^= t ^ (t >> 7)
	return b
}

// FlipDiagonalA8H1 returns a new bitboard flipped about the diagonal a8-h1
func (b Bitboard) FlipDiagonalA8H1() Bitboard {
	var t Bitboard
	k1 := Bitboard(0xaa00aa00aa00aa00)
	k2 := Bitboard(0xcccc0000cccc0000)
	k4 := Bitboard(0xf0f0f0f00f0f0f0f)
	t = b ^ (b << 36)
	b ^= k4 & (t ^ (b >> 36))
	t = k2 & (b ^ (b << 18))
	b ^= t ^ (t >> 18)
	t = k1 & (b ^ (b << 9))
	b ^= t ^ (t >> 9)
	return b
}

// Rotate180 returns a new bitboard rotated 180 degrees.
func (b Bitboard) Rotate180() Bitboard {
	return b.FlipVertical().FlipHorizontal()
}

// Rotate90 returns a bitboard rotated by 90 degrees (clockwise).
func (b Bitboard) Rotate90() Bitboard {
	return b.FlipDiagonalA1H8().FlipVertical()
}

// Rotate270 returns a bitboard rotated by 270 degrees (90 degrees counter-clockwise).
func (b Bitboard) Rotate270() Bitboard {
	return b.FlipVertical().FlipDiagonalA1H8()
}

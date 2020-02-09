package chess

import (
	"fmt"
	"strconv"
)

//-----------------------------------------------------------------------------
// Coordinate conversions
//-----------------------------------------------------------------------------

// AlgebraicToCartesian converts coordinates in algebraic notation to Cartesian coordinates.
func AlgebraicToCartesian(p string) (int, int) {
	symbols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var x int
	for i, v := range symbols {
		if string(p[0]) == v {
			x = i
		}
	}
	y, _ := strconv.Atoi(string(p[1]))
	return x, (y - 1)
}

// AlgebraicToBit converts coordinates in algebraic notation to an integer bit position.
func AlgebraicToBit(p string) int {
	x, y := AlgebraicToCartesian(p)
	return CartesianToBit(x, y)
}

// BitToAlgebraic converts an integer bit position to coordiantes in algebraic notation.
func BitToAlgebraic(p int) string {
	x, y := BitToCartesian(p)
	return CartesianToAlgebraic(x, y)
}

// BitToCartesian converts an integer bit position to Cartesian coordinates.
func BitToCartesian(p int) (int, int) {
	x := p % int(FILES)
	y := p / int(FILES)
	return x, y
}

// CartesianToAlgebraic converts Cartesian coordinates to coordinates in algebraic notation.
func CartesianToAlgebraic(x int, y int) string {
	symbols := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return fmt.Sprintf("%v%v", symbols[x], y+1)
}

// CartesianToBit converts Cartesian coordinates to an integer bit position.
func CartesianToBit(x int, y int) int {
	bit := y*int(FILES) + x
	return bit
}

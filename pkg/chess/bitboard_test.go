package chess

import (
	"fmt"
	"testing"
)

/*
letterR, rVertical, rHorizontal, rA1H8, and rA8H1 make the follow arrangements, respectively:
. 1 1 1 1 . . .
. 1 . . . 1 . .
. 1 . . . 1 . .
. 1 . . 1 . . .
. 1 1 1 . . . .
. 1 . 1 . . . .
. 1 . . 1 . . .
. 1 . . . 1 . .

. 1 . . . 1 . .   . . . 1 1 1 1 .   . . . . . . . .   . . . . . . . .
. 1 . . 1 . . .   . . 1 . . . 1 .   . . . . . . . .   1 1 1 1 1 1 1 1
. 1 . 1 . . . .   . . 1 . . . 1 .   1 . . . . 1 1 .   1 . . . 1 . . .
. 1 1 1 . . . .   . . . 1 . . 1 .   . 1 . . 1 . . 1   1 . . . 1 1 . .
. 1 . . 1 . . .   . . . . 1 1 1 .   . . 1 1 . . . 1   1 . . 1 . . 1 .
. 1 . . . 1 . .   . . . . 1 . 1 .   . . . 1 . . . 1   . 1 1 . . . . 1
. 1 . . . 1 . .   . . . 1 . . 1 .   1 1 1 1 1 1 1 1   . . . . . . . .
. 1 1 1 1 . . .   . . 1 . . . 1 .   . . . . . . . .   . . . . . . . .
*/
var (
	emptyBoard   Bitboard = Bitboard(0x0000000000000000)
	fullBoard    Bitboard = Bitboard(0xffffffffffffffff)
	darkSquares  Bitboard = Bitboard(0xaa55aa55aa55aa55)
	lightSquares Bitboard = Bitboard(0x55aa55aa55aa55aa)
	letterR      Bitboard = Bitboard(0x1e2222120e0a1222)
	rA1H8        Bitboard = Bitboard(0x000061928c88ff00)
	rA8H1        Bitboard = Bitboard(0x00ff113149860000)
	rVertical    Bitboard = Bitboard(0x22120a0e1222221e)
	rHorizontal  Bitboard = Bitboard(0x7844444870504844)
)

type params []interface{}

type action struct {
	Action     string
	ShouldPass bool
	Board      *Bitboard
	Params     params
}

func do(action string, b *Bitboard, indices ...int) {
	switch action {
	case "set":
		b.SetBit(indices[0])
	case "clear":
		b.ClearBit(indices[0])
	case "toggle":
		b.ToggleBit(indices[0])
	case "flipVertical":
		*b = b.FlipVertical()
	case "flipHorizontal":
		*b = b.FlipHorizontal()
	case "flipA1H8":
		*b = b.FlipDiagonalA1H8()
	case "flipA8H1":
		*b = b.FlipDiagonalA8H1()
	case "rotate180":
		*b = b.Rotate180()
	case "rotate90":
		*b = b.Rotate90()
	case "rotate270":
		*b = b.Rotate270()
	default:
		panic("invalid action")
	}
}

func check(action string, b Bitboard, params ...interface{}) (bool, error) {
	var err error
	switch action {
	case "compare":
		if expected, ok := params[0].(Bitboard); ok {
			if b != expected {
				err = fmt.Errorf("expected bitboard 0x%016x, actual: 0x%016x", expected, b)
			}
		} else {
			err = fmt.Errorf("invalid expected bitboard: %v", params[0])
		}
	case "get":
		if index, ok := params[0].(int); ok {
			if expected, ok := params[1].(int); ok {
				actual := b.GetBit(index)
				if expected != actual {
					err = fmt.Errorf("expected bit value %d, actual value: %d", expected, actual)
				}
			} else {
				err = fmt.Errorf("invalid expected bit: %v", params[1])
			}
		} else {
			err = fmt.Errorf("invalid bitboard index: %v", params[0])
		}
	case "isbitset":
		if index, ok := params[0].(int); ok {
			if expected, ok := params[1].(bool); ok {
				actual := b.IsBitSet(index)
				if expected != actual {
					err = fmt.Errorf("expected result %t, actual result: %t", expected, actual)
				}
			} else {
				err = fmt.Errorf("invalid expected result: %v", params[1])
			}
		} else {
			err = fmt.Errorf("invalid bitboard index: %v", params[0])
		}
	case "union":
		var boards []Bitboard
		for _, param := range params {
			if bitboard, ok := param.(Bitboard); ok {
				boards = append(boards, bitboard)
			}
		}
		expected := b
		actual := Union(boards...)
		if expected != actual {
			err = fmt.Errorf("expected board: %016x, actual board: %016x", expected, actual)
		}
	case "population":
		if expected, ok := params[0].(int); ok {
			actual := b.Population()
			if expected != actual {
				err = fmt.Errorf("expected population: %d, actual population: %d", expected, actual)
			}
		} else {
			err = fmt.Errorf("invalid requested population: %v", params[0])
		}
	default:
		err = fmt.Errorf("invalid action requested: %s", action)
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

func isCheck(action string) bool {
	for _, check := range []string{"compare", "get", "isbitset", "union", "population"} {
		if action == check {
			return true
		}
	}
	return false
}

func runTest(actions []action, t *testing.T) {
	for _, action := range actions {
		if isCheck(action.Action) {
			ok, err := check(action.Action, *action.Board, action.Params...)
			if ok != action.ShouldPass {
				if action.ShouldPass {
					t.Errorf("%s check failed: %s", action.Action, err)
				} else {
					t.Errorf("%s should have failed on bitboard 0x%016x", action.Action, *action.Board)
				}
			}
		} else {
			var indices []int
			for _, p := range action.Params {
				if index, ok := p.(int); ok {
					indices = append(indices, index)
				}
			}
			do(action.Action, action.Board, indices...)
		}
	}

}

func TestBitboardBitOperations(t *testing.T) {
	board := emptyBoard
	var allBoards params
	var actions []action
	for i := 0; i < 64; i++ {
		setBit := []action{
			action{"population", true, &board, params{0}},
			action{"set", true, &board, params{i}},
			action{"population", true, &board, params{1}},
			action{"get", true, &board, params{i, 1}},
			action{"isbitset", true, &board, params{i, true}},
			action{"compare", false, &board, params{emptyBoard}},
		}

		toggleBit := []action{
			action{"toggle", true, &board, params{i}},
			action{"compare", true, &board, params{emptyBoard}},
			action{"population", true, &board, params{0}},
		}

		setAndClearBit := []action{
			action{"set", true, &board, params{i}},
			action{"clear", true, &board, params{i}},
			action{"population", true, &board, params{0}},
		}
		actions = append(actions, setBit...)
		actions = append(actions, toggleBit...)
		actions = append(actions, setAndClearBit...)
	}

	for i := 0; i < 64; i++ {
		b := emptyBoard
		(&b).SetBit(i)
		allBoards = append(allBoards, b)
	}

	actions = append(actions, action{"union", true, &fullBoard, allBoards})

	runTest(actions, t)

}

func TestBitboardRotation(t *testing.T) {

	var tests []action
	board := lightSquares
	testRotate90 := []action{
		action{"rotate90", true, &board, params{}},
		action{"compare", true, &board, params{darkSquares}},
		action{"compare", false, &board, params{lightSquares}},
		action{"rotate90", true, &board, params{}},
		action{"compare", false, &board, params{darkSquares}},
		action{"compare", true, &board, params{lightSquares}},
		action{"flipA1H8", true, &board, params{}},
		action{"compare", true, &board, params{lightSquares}},
		action{"flipHorizontal", true, &board, params{}},
		action{"compare", false, &board, params{lightSquares}},
	}
	tests = append(tests, testRotate90...)

	board1, board2 := darkSquares, darkSquares
	testRotate180 := []action{
		action{"rotate180", true, &board1, params{}},
		action{"compare", true, &board1, params{darkSquares}},
		action{"flipVertical", true, &board2, params{}},
		action{"compare", false, &board2, params{darkSquares}},
		action{"compare", true, &board2, params{lightSquares}},
		action{"flipHorizontal", true, &board2, params{}},
		action{"compare", true, &board2, params{darkSquares}},
	}
	tests = append(tests, testRotate180...)

	board3 := lightSquares
	testRotate270 := []action{
		action{"rotate270", true, &board3, params{}},
		action{"compare", false, &board3, params{lightSquares}},
		action{"compare", true, &board3, params{darkSquares}},
		action{"rotate90", true, &board3, params{}},
		action{"compare", true, &board3, params{lightSquares}},
	}
	tests = append(tests, testRotate270...)

	r1, r2, r3, r4 := letterR, letterR, letterR, letterR
	testFlips := []action{
		action{"population", true, &r1, params{19}},
		action{"flipA1H8", true, &r1, params{}},
		action{"compare", true, &r1, params{rA1H8}},
		action{"population", true, &r1, params{19}},
		action{"population", true, &r2, params{19}},
		action{"flipA8H1", true, &r2, params{}},
		action{"compare", true, &r2, params{rA8H1}},
		action{"population", true, &r2, params{19}},
		action{"population", true, &r3, params{19}},
		action{"flipVertical", true, &r3, params{}},
		action{"compare", true, &r3, params{rVertical}},
		action{"population", true, &r3, params{19}},
		action{"population", true, &r4, params{19}},
		action{"flipHorizontal", true, &r4, params{}},
		action{"compare", true, &r4, params{rHorizontal}},
		action{"population", true, &r4, params{19}},
	}
	tests = append(tests, testFlips...)
	runTest(tests, t)

}

package main

import (
	"github.com/stephensli/advent-of-code-2021/helpers"
)

type SeaCucumber int

const (
	Blank       SeaCucumber = 0
	EastFacing  SeaCucumber = 1
	SouthFacing SeaCucumber = 2
)

func (s SeaCucumber) Empty() bool {
	return s == Blank
}

func (s SeaCucumber) IsEastFacing() bool {
	return s == EastFacing
}

func (s SeaCucumber) IsSouthFacing() bool {
	return s == SouthFacing
}

func StepCucumbers(input [][]SeaCucumber) ([][]SeaCucumber, int) {
	// east facing is actioned first, so lets go through all the rows and check
	// to see if we can move any. One can only move if its right most position
	// is empty.
	eastMoveResult := helpers.Clone(input)
	changeCount := 0

	for column := 0; column < len(input); column++ {
		for row := 0; row < len(input[column]); row++ {
			if input[column][row] != EastFacing {
				continue
			}

			// if we are not the last part and the next space is free, then lets go and
			// move our sea cucumber over to the next part, otherwise we will check the
			// edge.
			if row+1 < len(input[column]) && input[column][row+1] == Blank {
				eastMoveResult[column][row] = Blank
				eastMoveResult[column][row+1] = EastFacing

				changeCount += 1
				continue
			}

			if row+1 == len(input[column]) && input[column][0] == Blank {
				eastMoveResult[column][row] = Blank
				eastMoveResult[column][0] = EastFacing

				changeCount += 1
				continue
			}

		}
	}

	southMoveResult := helpers.Clone(eastMoveResult)

	// move down instead of left to right.
	for column := 0; column < len(input[0]); column++ {
		for row := 0; row < len(input); row++ {
			if eastMoveResult[row][column] != SouthFacing {
				continue
			}

			// if we are not the last part and the next space is free, then lets go and
			// move our sea cucumber over to the next part, otherwise we will check the
			// edge.
			if row+1 < len(eastMoveResult) && eastMoveResult[row+1][column] == Blank {
				southMoveResult[row][column] = Blank
				southMoveResult[row+1][column] = SouthFacing

				changeCount += 1
				continue
			}

			if row+1 == len(eastMoveResult) && eastMoveResult[0][column] == Blank {
				southMoveResult[row][column] = Blank
				southMoveResult[0][column] = SouthFacing

				changeCount += 1
				continue
			}
		}
	}

	return southMoveResult, changeCount
}

package main

import (
	"fmt"
)

type Board struct {
	content          [5][5]int
	matchedValues    map[int]bool
	lastUpdatedValue int
}

func (b *Board) hasCompletedRowOrColumn() bool {
	// check rows
	for i := 0; i < 5; i++ {
		allValid := true

		for j := 0; j < 5; j++ {
			val := b.content[i][j]

			if in, ok := b.matchedValues[val]; !ok || !in {
				allValid = false
				break
			}
		}

		if allValid {
			return true
		}
	}

	// columns columns
	for j := 0; j < 5; j++ {
		allValid := true

		for i := 0; i < 5; i++ {
			val := b.content[i][j]

			if in, ok := b.matchedValues[val]; !ok || !in {
				allValid = false
				break
			}
		}

		if allValid {
			return true
		}
	}
	return false
}

func (b *Board) addMatchedPosition(value int) (hasCompletedRowOrColumn bool) {
	for i := 0; i < len(b.content); i++ {
		for j := 0; j < len(b.content[i]); j++ {
			if b.content[i][j] == value {
				b.matchedValues[value] = true
				b.lastUpdatedValue = value
				return b.hasCompletedRowOrColumn()
			}
		}
	}

	return b.hasCompletedRowOrColumn()
}

func (b *Board) getSumTotalUnmatchedPositions() int {
	sum := 0

	for i := 0; i < len(b.content); i++ {
		for j := 0; j < len(b.content[i]); j++ {
			if val, ok := b.matchedValues[b.content[i][j]]; !ok || !val {
				sum += b.content[i][j]
			}
		}
	}

	return sum
}

func first(source []int, boards []*Board) int {
	for _, val := range source {
		for _, board := range boards {
			if board.addMatchedPosition(val) {
				return board.lastUpdatedValue * board.getSumTotalUnmatchedPositions()
			}
		}
	}

	return -1
}

func second(source []int, boards []*Board) int {
	won := map[int]bool{}
	lastWonBoard := 0

	for _, val := range source {
		for boardI, board := range boards {
			if vl, ok := won[boardI]; ok && vl {
				continue
			}

			if board.addMatchedPosition(val) {
				won[boardI] = true
				lastWonBoard = boardI
			}
		}
	}

	return boards[lastWonBoard].lastUpdatedValue *
		boards[lastWonBoard].getSumTotalUnmatchedPositions()
}

func main() {
	source, boards := readFileInput("./2021/day4/input.txt")
	fmt.Println("first", first(source, boards))

	source, boards = readFileInput("./2021/day4/input.txt")
	fmt.Println("second", second(source, boards))
}

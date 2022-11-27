package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"math"
)

func determineFloor(input []string) (int int, basementIndex int) {
	basementIndex = math.MinInt
	floor := 0

	for i, value := range input {
		switch value {
		case ")":
			floor -= 1
			break
		case "(":
			floor += 1
			break
		}

		if basementIndex == math.MinInt && floor < 0 {
			basementIndex = i
		}

	}

	if basementIndex != math.MinInt {
		basementIndex += 1
	}

	return floor, basementIndex
}

func main() {
	path, deferFunc := aoc.Setup(2015, 1, false)
	defer deferFunc()

	input := file.ToTextSplit(path, "")
	targetFloor, basementIndex := determineFloor(input[0])

	aoc.PrintAnswer(1, targetFloor)
	aoc.PrintAnswer(2, basementIndex)
}

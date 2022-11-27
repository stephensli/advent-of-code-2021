package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"github.com/stephensli/aoc/helpers/printers"
)

func parseInput(lines [][]string) [][]SeaCucumber {
	grid := make([][]SeaCucumber, len(lines))

	for i, line := range lines {
		row := make([]SeaCucumber, len(line))

		for j, value := range line {
			switch value {
			case ".":
				row[j] = Blank
				break
			case ">":
				row[j] = EastFacing
				break
			case "v":
				row[j] = SouthFacing
				break
			}

		}

		grid[i] = row
	}

	return grid
}

func partOne(input [][]SeaCucumber) {
	// every step, teh sea cucumbers in the EAST facing herd attempt to move forward one location.
	// then teh sea cucumbers' int the south facing herd attempt to move one location.
	//
	// when a herd moves forward, every sea cucumber in the herd first simultaneous considers whether there
	// is a sea cucumber in the adjacent location its facing. (even another sea cucumber facing the same direction).
	//
	// then every sea cucumber facing an empty location simultaneously moves into that location.
	//
	// Due to strong water currents in the area, sea cucumbers that move off the right edge of the map appear on the
	// left edge, and sea cucumbers that move off the bottom edge of the map appear on the top edge.
	resp, changes := StepCucumbers(input)

	count := 1
	for {
		if changes == 0 {
			break
		}

		count += 1

		resp, changes = StepCucumbers(resp)
	}

	aoc.PrintAnswer(1, count)
}

func main() {
	path, deferFunc := aoc.Setup(2021, 25, false)
	defer deferFunc()

	lines := file.ToTextSplit(path, "")
	input := parseInput(lines)
	printers.JsonPrint(input, true)

	partOne(input)
}

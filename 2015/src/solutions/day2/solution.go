package main

import (
	"fmt"
	"github.com/life4/genesis/slices"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

func parseInput(input []string) []Present {
	return slices.Map(input, func(element string) Present {
		var length, width, height int

		_, _ = fmt.Sscanf(element, "%dx%dx%d", &length, &width, &height)
		return Present{Dimensions{length, width, height}}
	})
}

func partOne(prisms []Present) int {
	return slices.Reduce(prisms, 0, func(elm Present, acc int) int {
		return elm.GetWrappingPaperWithSlack() + acc
	})
}

func partTwo(prisms []Present) int {
	return slices.Reduce(prisms, 0, func(elm Present, acc int) int {
		return elm.GetRibbonWithBow() + acc
	})
}

func main() {
	path, deferFunc := aoc.Setup(2015, 2, false)
	defer deferFunc()

	input := file.ToTextLines(path)
	prisms := parseInput(input)

	aoc.PrintAnswer(1, partOne(prisms))
	aoc.PrintAnswer(2, partTwo(prisms))
}

package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers"
	"math"
)

func main() {
	// nums := strings.Split(helpers.ReadFileToTextLines("./day7/input.txt")[0], ",")
	nums := helpers.ReadFileToNumbersSplit("./day7/input.txt", ",")[0]

	first := false

	min := -1
	max := -1

	for _, num := range nums {
		if min == -1 {
			min = num
			max = num
		}

		if num < min {
			min = num
		}

		if num > max {
			max = num
		}

	}

	lowestCost := -1

	for alignmentNumber := min; alignmentNumber < max; alignmentNumber++ {
		cost := 0

		for _, num := range nums {
			baseCost := int(math.Abs(float64(alignmentNumber - num)))

			if first {
				cost += baseCost
			} else {
				cost += (baseCost * (baseCost + 1)) / 2

			}

			if cost > lowestCost && lowestCost != -1 {
				break
			}
		}

		if cost < lowestCost || lowestCost == -1 {
			lowestCost = cost
		}
	}

	fmt.Println(lowestCost)
}

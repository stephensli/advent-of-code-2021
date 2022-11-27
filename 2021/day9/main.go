package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers"
	"sort"
)

type lowestPoint struct {
	value     int
	basinSize int
}

func determineBasinSize(input [][]int, i, j, count int, history map[string]bool) int {
	history[fmt.Sprintf("%d%d", i, j)] = true

	if input[i][j] == 9 {
		return count
	}

	// EDGE CASE
	// Everything greater than or equal to the value can flow down.
	// When it says "flow" it means not in increments of 1, but any
	// number greater than the value.

	value := input[i][j]
	fmt.Printf("i %d, j %d - value: %v\n", i, j, value)

	// check top
	if val, ok := history[fmt.Sprintf("%d%d", i-1, j)]; !ok || !val {
		if i != 0 && input[i-1][j] >= value {
			count = determineBasinSize(input, i-1, j, count, history)
		}
	}

	//	 check left
	if val, ok := history[fmt.Sprintf("%d%d", i, j-1)]; !ok || !val {
		if j != 0 && input[i][j-1] >= value {
			count = determineBasinSize(input, i, j-1, count, history)
		}
	}

	// check right
	if val, ok := history[fmt.Sprintf("%d%d", i, j+1)]; !ok || !val {
		if j != len(input[i])-1 && input[i][j+1] >= value {
			count = determineBasinSize(input, i, j+1, count, history)
		}
	}

	// check bottom
	if val, ok := history[fmt.Sprintf("%d%d", i+1, j)]; !ok || !val {
		if i != len(input)-1 && input[i+1][j] >= value {
			count = determineBasinSize(input, i+1, j, count, history)
		}
	}

	return count + 1
}

func getLowPoints(input [][]int) []lowestPoint {
	output := []lowestPoint{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {

			// check top
			if i != 0 && input[i-1][j] <= input[i][j] {
				continue
			}

			// check left
			if j != 0 && input[i][j-1] <= input[i][j] {
				continue
			}

			// check right
			if j != len(input[i])-1 && input[i][j+1] <= input[i][j] {
				continue
			}

			// check bottom
			if i != len(input)-1 && input[i+1][j] <= input[i][j] {
				continue
			}

			fmt.Println("====")

			history := map[string]bool{}
			output = append(output, lowestPoint{
				value:     input[i][j],
				basinSize: determineBasinSize(input, i, j, 0, history),
			})

		}
	}

	return output

}

func main() {
	mapping := helpers.ReadFileToNumbersSplit("./day9/input.txt", "")

	lowestPoints := getLowPoints(mapping)
	sum := 0

	for _, v := range lowestPoints {
		sum += v.value + 1
	}

	sort.Slice(lowestPoints, func(i, j int) bool {
		return lowestPoints[i].basinSize > lowestPoints[j].basinSize
	})

	fmt.Printf("lowestPoints: %v\n", lowestPoints)
	output := lowestPoints[0].basinSize * lowestPoints[1].basinSize * lowestPoints[2].basinSize

	fmt.Printf("first: %v\n", sum)
	fmt.Printf("second: %v\n", output)

}

package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers"
)

type position struct {
	x, y int
}

type Octopus struct {
	value      int
	hasFlashed bool
}

func parseInput(lines [][]int) [][]Octopus {
	var grid [][]Octopus

	for _, v := range lines {
		var entry []Octopus

		for _, val := range v {
			entry = append(entry, Octopus{val, false})
		}

		grid = append(grid, entry)
	}

	return grid
}

func printGrid(grid [][]Octopus) {
	fmt.Println("---")

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			fmt.Print(grid[i][j].value, ",")
		}
		fmt.Println()
	}

}

func flashOctopuses(grid [][]Octopus, positions []position) int {
	allFlashedPositions := make([]position, len(positions))
	copy(allFlashedPositions, positions)

	for {
		nextPositions := []position{}

		for _, p := range positions {
			if grid[p.x][p.y].hasFlashed {
				continue
			}
			// increase all positions
			posToCheck := []position{
				{x: p.x - 1, y: p.y},
				{x: p.x + 1, y: p.y},
				{x: p.x, y: p.y - 1},
				{x: p.x, y: p.y + 1},
				{x: p.x - 1, y: p.y - 1},
				{x: p.x + 1, y: p.y - 1},
				{x: p.x - 1, y: p.y + 1},
				{x: p.x + 1, y: p.y + 1},
			}

			for _, posToCheck := range posToCheck {
				if posToCheck.x < 0 || posToCheck.x > len(grid)-1 || posToCheck.y < 0 || posToCheck.y > len(grid[posToCheck.x])-1 {
					continue
				}

				if grid[posToCheck.x][posToCheck.y].hasFlashed {
					continue
				}

				grid[posToCheck.x][posToCheck.y].value += 1

				if grid[posToCheck.x][posToCheck.y].value > 9 {
					nextPositions = append(nextPositions, position{posToCheck.x, posToCheck.y})
				}
			}

			grid[p.x][p.y].hasFlashed = true
			allFlashedPositions = append(allFlashedPositions, nextPositions...)
		}

		if len(nextPositions) == 0 {
			break
		}

		positions = nextPositions

	}

	totalFlash := 0

	for _, p := range allFlashedPositions {

		if grid[p.x][p.y].hasFlashed {
			grid[p.x][p.y].value = 0
			grid[p.x][p.y].hasFlashed = false

			// this is because duplicates can exist.
			totalFlash += 1
		}

	}

	return totalFlash
}

func tickAllByValue(grid [][]Octopus, amount int) (flashPositions []position) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j].value += 1

			if grid[i][j].value > 9 {
				flashPositions = append(flashPositions, position{i, j})
			}
		}
	}

	return flashPositions
}

func processDay(grid [][]Octopus) int {
	totalFlash := 0

	for {
		pos := tickAllByValue(grid, 1)

		if len(pos) == 0 {
			break
		}

		totalFlash += flashOctopuses(grid, pos)
		break
	}

	//printGrid(grid)

	return totalFlash
}

func main() {
	lines := helpers.ReadFileToNumbersSplit("./day11/input.txt", "")
	grid := parseInput(lines)

	totalFlash := 0

	i := 0

	completed := 0

	for {
		amount := processDay(grid)

		if i == 100 {
			fmt.Printf("first: %v\n", totalFlash)
			completed += 1
		}

		if amount == 100 {
			fmt.Println("all flashed on day ", i+1)
			completed += 1
		}

		if completed == 2 {
			break
		}

		totalFlash += amount
		i += 1
	}

}

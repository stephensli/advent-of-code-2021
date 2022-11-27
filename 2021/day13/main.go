package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stephensli/aoc/helpers"
)

type Point struct {
	x int `json:x`
	y int `json:y`
}

func parseInput(lines []string) (grid [][]string, folds []Point) {
	maxY := 0
	maxX := 0

	var foldLines []string
	gridValues := []Point{}

	// first go and parse the grid values and determine the min
	// max which will be used to determine the grid size.
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			foldLines = lines[i+1:]
			break
		}

		splitLine := strings.Split(line, ",")

		y, _ := strconv.Atoi(splitLine[0])
		x, _ := strconv.Atoi(splitLine[1])

		gridValues = append(gridValues, Point{x, y})

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	grid = make([][]string, maxX+1)

	for i := 0; i < maxX+1; i++ {
		grid[i] = make([]string, maxY+1)
		for j := 0; j < maxY+1; j++ {
			grid[i][j] = "."
		}
	}

	for _, gridVal := range gridValues {
		grid[gridVal.x][gridVal.y] = "#"

	}

	folds = []Point{}

	for _, v := range foldLines {
		fold := strings.Split(string(v[11:]), "=")
		val, _ := strconv.Atoi(fold[1])

		if fold[0] == "x" {
			folds = append(folds, Point{val, -1})
		} else {
			folds = append(folds, Point{-1, val})
		}
	}

	return grid, folds
}

func foldGridX(grid [][]string, fold int) [][]string {
	newGrid := make([][]string, len(grid))

	for i := range newGrid {
		newGrid[i] = make([]string, fold)

		for j := range newGrid[i] {
			newGrid[i][j] = grid[i][j]
		}
	}

	for i, v := range grid {
		grid[i] = v[fold+1:]
	}

	for i := range grid {
		for j := range grid[i] {
			if newGrid[i][fold-j-1] != "#" {
				newGrid[i][fold-j-1] = grid[i][j]
			}
		}
	}

	return newGrid
}

func foldGridY(grid [][]string, fold int) [][]string {
	newGrid := make([][]string, fold)

	for i := range newGrid {
		newGrid[i] = make([]string, len(grid[0]))

		for j := range newGrid[i] {
			newGrid[i][j] = grid[i][j]
		}
	}

	grid = grid[fold+1:]

	for i := range grid {
		for j := range grid[i] {
			if newGrid[fold-i-1][j] != "#" {
				newGrid[fold-i-1][j] = grid[i][j]
			}
		}
	}

	return newGrid
}

func foldGrid(grid [][]string, fold Point) (outGrid [][]string) {
	if fold.x != -1 {
		return foldGridX(grid, fold.x)
	}

	return foldGridY(grid, fold.y)
}

func main() {
	lines := helpers.ReadFileToTextLines("./day13/input.txt")
	grid, folds := parseInput(lines)

	for _, p := range folds {
		grid = foldGrid(grid, p)
	}

	helpers.JsonPrint(grid, "day13", true)
	count := 0

	fmt.Println("second:")
	for i, _ := range grid {
		for _, v := range grid[i] {

			if v == "#" {
				fmt.Print(v)
				count += 1
			} else {

				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Printf("first: %v\n", count)

}

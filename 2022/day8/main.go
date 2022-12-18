package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

// First, determine whether there is enough tree cover here to keep a tree house
// hidden. To do this, you need to count the number of trees that are visible
// from outside the grid when looking directly along a row or column.
//
// Each tree is represented as a single digit whose value is its height, where
// 0 is the shortest and 9 is the tallest.
//
// A tree is visible if all of the other trees between it and an edge of the
// grid are shorter than it. Only consider trees in the same row or column;
// that is, only look up, down, left, or right from any given tree.
//
// All of the trees around the edge of the grid are visible - since they are
// already on the edge, there are no trees to block the view. In this example,
// that only leaves the interior nine trees to consider:
//
// 30373
// 25512
// 65332
// 33549
// 35390
//
// - The top-left 5 is visible from the left and top. (It isn't visible from
// the right or bottom since other trees of height 5 are in the way.)
//
// - The top-middle 5 is visible from the top and right.
//
// - The top-right 1 is not visible from any direction; for it to be visible,
// there would need to only be trees of height 0 between it and an edge.
//
// - The left-middle 5 is visible, but only from the right.
//
// - The center 3 is not visible from any direction; for it to be visible,
// there would need to be only trees of at most height 2 between it and an edge.
//
// - The right-middle 3 is visible from the right.
//
// - In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
//
// With 16 trees visible on the edge and another 5 visible in the interior,
// a total of 21 trees are visible in this arrangement.
//
// Consider your map; how many trees are visible from outside the grid?
//
// **A tree is visible if all of the other trees between it and an edge of the
// grid are shorter than it**
func main() {
	path, complete := aoc.Setup(2022, 8, false)
	defer complete()

	grid := file.ToNumbersSplit(path, "")

	var first int
	var second int

	topSize := len(grid)*2 - 4
	sideSize := len(grid[0]) * 2

	first += topSize + sideSize

	// loop the inner grid, not including the outer layer as this layer is
	// always visible.
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			visible, dist := isVisible(grid, i, j)

			if visible {
				first += 1
			}

			if dist > second {
				second = dist
			}
		}
	}

	aoc.PrintAnswer(1, first)
	aoc.PrintAnswer(2, second)
}

func isVisible(grid [][]int, i, j int) (bool, int) {
	left, distLeft := checkLeft(grid, grid[i][j], i, j, 0)
	right, distRight := checkRight(grid, grid[i][j], i, j, 0)
	up, distTop := checkUp(grid, grid[i][j], i, j, 0)
	down, distDown := checkDown(grid, grid[i][j], i, j, 0)

	return left || right || up || down, distLeft * distRight * distTop * distDown

}

package algorithms

import (
	"strconv"
)

func GetGridNeighbors(grid [][]Node, directions []Direction, source Coords) []Node {
	neighbors := []Node{}

	for _, d := range directions {
		x, y := source.X, source.Y

		switch d {
		case UpDirection:
			if source.X-1 >= 0 {
				neighbors = append(neighbors, grid[x-1][y])
				continue
			}
		case LeftDirection:
			if source.Y-1 >= 0 {
				neighbors = append(neighbors, grid[x][y-1])
				continue
			}
		case RightDirection:
			if source.Y+1 < len(grid[source.X]) {
				neighbors = append(neighbors, grid[x][y+1])
				continue
			}
		case DownDirection:
			if source.X+1 < len(grid) {
				neighbors = append(neighbors, grid[x+1][y])
				continue
			}
		default:
			panic("direction not implemented: " + strconv.Itoa(int(d)))
		}
	}

	return neighbors
}

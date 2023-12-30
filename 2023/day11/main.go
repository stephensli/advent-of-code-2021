package main

import (
	"fmt"
	"sync"

	"github.com/stephensli/aoc/helpers/algorithms"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

func parse(grid [][]string) [][]Position {
	positions := make([][]Position, len(grid))

	galexyIndex := 1

	for i := 0; i < len(grid); i++ {
		row := make([]Position, len(grid[i]))
		positions[i] = row

		for j := 0; j < len(grid[i]); j++ {
			position := Position{
				PositionType:  EmptySpaceType,
				PositionValue: 0,
				SpaceWeight:   1,
			}

			if grid[i][j] == "#" {
				position.PositionType = GalexySpaceType
				position.PositionValue = galexyIndex
				galexyIndex += 1
			}

			row[j] = position
		}
	}

	return positions
}

func expandSpace(p1 bool, grid [][]Position) [][]Position {
	spaceWeight := 1

	if !p1 {
		spaceWeight = 999999
	}

	// rows
	emptyRow := []Position{}
	for i := 0; i < len(grid); i++ {
		emptyRow = append(emptyRow, Position{
			SpaceWeight:   spaceWeight,
			Coords:        algorithms.Coords{},
			PositionType:  0,
			PositionValue: 0,
		})
	}

	for i := 0; i < len(grid); i++ {
		empty := true

		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].Galexy() {
				empty = false
				break
			}
		}

		if empty {
			grid = append(grid[:i+1], grid[i:]...)
			grid[i] = emptyRow
			i += 1
		}
	}

	// columns
	for i := 0; i < len(grid[0]); i++ {
		empty := true

		for j := 0; j < len(grid); j++ {
			if grid[j][i].Galexy() {
				empty = false
				break
			}
		}

		if empty {
			for j := 0; j < len(grid); j++ {
				grid[j] = append(grid[j][:i+1], grid[j][i:]...)
				grid[j][i] = Position{
					SpaceWeight:   spaceWeight,
					Coords:        algorithms.Coords{},
					PositionType:  0,
					PositionValue: 0,
				}
			}

			i += 1
		}
	}

	return grid
}

func getGalexyPositions(grid [][]Position) ([]algorithms.Coords, map[int]algorithms.Coords) {
	list := []algorithms.Coords{}
	mapView := map[int]algorithms.Coords{}

	for i, v := range grid {
		for j, p := range v {
			if p.Galexy() {
				list = append(list, algorithms.Coords{X: i, Y: j})
				mapView[p.PositionValue] = algorithms.Coords{X: i, Y: j}
			}
		}
	}

	return list, mapView
}

func main() {
	isPartOne := false

	path, complete := aoc.Setup(2023, 11, false)
	defer complete()

	grid := expandSpace(isPartOne, parse(file.ToTextSplit(path, "")))
	galexies, _ := getGalexyPositions(grid)

	outcome := 0
	history := map[string]bool{}

	var nodes [][]algorithms.Node

	for i := 0; i < len(grid); i++ {
		var innerNodes []algorithms.Node
		for j := 0; j < len(grid[i]); j++ {
			p := grid[i][j]

			p.Coords = algorithms.Coords{X: i, Y: j}
			innerNodes = append(innerNodes, p)
		}
		nodes = append(nodes, innerNodes)
	}

	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	for i, c := range galexies {
		wg.Add(1)

		i := i
		c := c

		go func() {
			_, distance := algorithms.Bfs(nodes, algorithms.NonDigagnonalDirections, c, c)
			defer wg.Done()

			for j, c2 := range galexies {
				historykeyOne := fmt.Sprintf("%d:%d:%d:%d", c.X, c.Y, c2.X, c2.Y)
				historykeyTwo := fmt.Sprintf("%d:%d:%d:%d", c2.X, c2.Y, c.X, c.Y)

				mx.Lock()

				if i != j && !history[historykeyOne] && !history[historykeyTwo] {
					history[historykeyOne] = true
					history[historykeyTwo] = true
					outcome += distance[c2]
				}
				mx.Unlock()
			}
		}()

	}
	wg.Wait()

	aoc.PrintAnswerAny(outcome)
}


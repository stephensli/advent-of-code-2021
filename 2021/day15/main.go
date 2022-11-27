package main

import (
	"container/heap"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"github.com/stephensli/aoc/helpers/queue"
	"math"
)

type Position struct {
	x, y int
}

func getNextValue(grid [][]int, size, x, y int) int {
	shiftX := x / size
	shiftY := y / size

	orgX := x - shiftX*size
	orgY := y - shiftY*size

	shiftValue := grid[orgY][orgX] + shiftX + shiftY

	if shiftValue > 9 {
		shiftValue = int(math.Mod(float64(shiftValue), 9))

	}

	return shiftValue
}

func Dijkstra(grid [][]int, growth int) int {
	width := len(grid[0]) * growth
	height := len(grid) * growth

	costs := make([][]int, height)
	for i := range costs {
		costs[i] = make([]int, width)

		for i2 := range costs[i] {
			costs[i][i2] = math.MaxInt
		}
	}

	costs[0][0] = 0

	pq := queue.MinPriorityQueue{{
		Value:    Position{0, 0},
		Priority: 0,
		Index:    0,
	},
	}
	heap.Init(&pq)

	for {
		if pq.Len() == 0 {
			break
		}

		popped := heap.Pop(&pq).(*queue.Item)

		pv := popped.Value.(Position)
		x, y := pv.x, pv.y

		if x == width-1 && y == height-1 {
			continue
		}

		posToCheck := []Position{
			{x: x + 1, y: y},
			{x: x, y: y + 1},
			{x: x, y: y - 1},
			{x: x - 1, y: y},
		}

		for _, p := range posToCheck {
			if p.x < 0 || p.x > len(costs)-1 || p.y < 0 || p.y > len(costs[p.x])-1 {
				continue
			}

			ny := p.y
			nx := p.x

			nextValue := getNextValue(grid, len(grid), nx, ny)
			nextCost := costs[y][x] + nextValue

			if nextCost < costs[ny][nx] {
				costs[ny][nx] = nextCost
				heap.Push(&pq, &queue.Item{Value: p, Priority: nextCost})
			}
		}
	}

	return costs[height-1][width-1]
}

func main() {
	defer aoc.Setup(2021, 15)()

	lines := file.ToNumbersSplit("./input.txt", "")

	aoc.PrintAnswer(1, Dijkstra(lines, 1))
	aoc.PrintAnswer(2, Dijkstra(lines, 5))
}

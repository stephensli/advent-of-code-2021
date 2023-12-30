package algorithms

import (
	"github.com/stephensli/aoc/helpers/cache"
	"github.com/stephensli/aoc/helpers/queue"
)

func Bfs(grid [][]Node, directions []Direction, source, target Coords) (prev map[Coords]Coords, dist map[Coords]int) {
	previous := map[Coords]Coords{}
	distance := map[Coords]int{}

	queue := queue.Queue[Node]{}
	queue.Push(grid[source.X][source.Y])

	gridCache := cache.New[Coords, bool]()
	gridCache.Set(source, true)
	distance[source] = 0

	for queue.Len() != 0 {
		queueValue := queue.Pop()

		for _, neighbor := range GetGridNeighbors(grid, directions, queueValue.Position()) {
			if gridCache.Has(neighbor.Position()) {
				continue
			}

			gridCache.Set(neighbor.Position(), true)

			distance[neighbor.Position()] = distance[queueValue.Position()] + neighbor.Weight()
			previous[neighbor.Position()] = queueValue.Position()
			queue.Push(neighbor)

			if neighbor.Position() == target {
				return previous, distance
			}
		}
	}

	return previous, distance
}

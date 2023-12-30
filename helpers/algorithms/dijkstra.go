package algorithms

import (
	"container/heap"
	"math"

	"github.com/stephensli/aoc/helpers/cache"
	"github.com/stephensli/aoc/helpers/queue"
)

// 1  function Dijkstra(Graph, source):
// 2
// 3      for each vertex v in Graph.Vertices:
// 4          dist[v] ← INFINITY
// 5          prev[v] ← UNDEFINED
// 6          add v to Q
// 7      dist[source] ← 0
// 8
// 9      while Q is not empty:
//10          u ← vertex in Q with min dist[u]
//11          remove u from Q
//12
//13          for each neighbor v of u still in Q:
//14              alt ← dist[u] + Graph.Edges(u, v)
//15              if alt < dist[v]:
//16                  dist[v] ← alt
//17                  prev[v] ← u
//18
//19      return dist[], prev[]

// A node is a single point on the grid in which the Dijkstra algorithm can use
// to determine if it can move forward onto that slot or not. It also includes
// additional supporting information like cost and if its a wall (blocking).
type Node interface {
	// Wall returns true if and only if the block cannot be stepped on. The
	// direction is the source direction the path has taken to hit said wall, e.g
	// if the direction was LeftDirection, the path came from the right.
	Wall(direction Direction) bool
	Position() Coords
	Weight() int
	Value() int
}

// DijkstraGrid computes the shortest distance in a grid between the source and
// the destination nodes. Using the allowed directions to determine how it can
// move around the grid.
//
// Implementation of the Node interface is required to determine if blocks are
// walls or not.
func DijkstraGrid(
	grid [][]Node,
	allowedDirections []Direction,
	source, target Coords) (shortestPath int, distance []Node, previous map[Coords]Coords) {
	previous = map[Coords]Coords{}

	minPriorityQueue := queue.MinPriorityQueue{{
		Value:    grid[source.X][source.Y],
		Priority: 0,
		Index:    0,
	}}

	gridCache := cache.New[Coords, bool]()
	heap.Init(&minPriorityQueue)

	shortestPath = math.MaxInt

	for len(minPriorityQueue) != 0 {
		queueItem := heap.Pop(&minPriorityQueue).(*queue.Item)
		node := queueItem.Value.(Node)

		if gridCache.Has(node.Position()) {
			continue
		}

		if node.Position() == target {
			shortestPath = queueItem.Priority
			next := node.Position()

			for next != source {
				distance = append(distance, grid[next.X][next.Y])
				next = previous[next]
			}

			break
		}

		gridCache.Set(node.Position(), true)

		for _, neighbor := range GetGridNeighbors(grid, allowedDirections, node.Position()) {
			if gridCache.Has(neighbor.Position()) {
				continue
			}

			previous[neighbor.Position()] = node.Position()
			heap.Push(&minPriorityQueue, &queue.Item{
				Value:    neighbor,
				Priority: queueItem.Priority + neighbor.Weight(),
			})
		}
	}

	distance = append(distance, grid[source.X][source.Y])
	for i, j := 0, len(distance)-1; i < j; i, j = i+1, j-1 {
		distance[i], distance[j] = distance[j], distance[i]
	}

	return
}

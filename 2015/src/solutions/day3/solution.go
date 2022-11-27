package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/cache"
	"github.com/stephensli/aoc/helpers/file"
)

type Position struct {
	X, Y int
}

func (p *Position) Move(input string) {
	switch input {
	case "^":
		// up
		p.Y -= 1
		break
	case "v":
		// down
		p.Y += 1
		break
	case ">":
		// right
		p.X += 1
		break
	case "<":
		// left
		p.X -= 1
		break
	}
}

func (p Position) Key() string {
	return fmt.Sprintf("%d/%d", p.X, p.Y)
}

func partOne(input []string) int {
	visited := cache.New[string, int]()

	// defaults to 0,0
	pos := Position{0, 0}
	visited.Set(pos.Key(), 1)

	for _, value := range input {
		pos.Move(value)

		if val, ok := visited.Get(pos.Key()); ok {
			visited.Set(pos.Key(), val+1)
		} else {
			visited.Set(pos.Key(), 1)
		}
	}

	return visited.Len()
}

func partTwo(input []string) int {
	visited := cache.New[string, int]()

	santa := &Position{0, 0}
	robot := &Position{0, 0}
	visited.Set(santa.Key(), 1)

	for i, value := range input {
		pos := santa

		if i%2 == 0 {
			pos = robot
		}

		pos.Move(value)

		if val, ok := visited.Get(pos.Key()); ok {
			visited.Set(pos.Key(), val+1)
		} else {
			visited.Set(pos.Key(), 1)
		}
	}

	return visited.Len()
}

func main() {
	path, deferFunc := aoc.Setup(2015, 3, false)
	defer deferFunc()

	input := file.ToTextSplit(path, "")[0]

	aoc.PrintAnswer(1, partOne(input))
	aoc.PrintAnswer(2, partTwo(input))
}

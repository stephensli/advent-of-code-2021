package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

func partOne(games []*Game) int64 {
	var idSum int64

GameLoop:
	for _, g := range games {
		for _, c := range g.cubes {
			if c.color == red && c.count > 12 {
				continue GameLoop
			}
			if c.color == green && c.count > 13 {
				continue GameLoop
			}
			if c.color == blue && c.count > 14 {
				continue GameLoop
			}
		}
		idSum += g.id
	}
	return idSum
}

func partTwo(games []*Game) int64 {
	var idSum int64

	for _, g := range games {
		minMap := map[Color]int64{}
		for _, c := range g.cubes {
			if value, ok := minMap[c.color]; !ok || c.count > value {
				minMap[c.color] = c.count
			}
		}

		idSum += (minMap[red] * minMap[blue] * minMap[green])
	}

	return idSum
}

func main() {
	path, complete := aoc.Setup(2023, 2, false)
	defer complete()

	lines := file.ToTextLines(path)
	games := parseInput(lines)

	aoc.PrintAnswer(1, partOne(games))
	aoc.PrintAnswer(2, partTwo(games))
}

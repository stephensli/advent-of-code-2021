package main

import (
	"sort"
	"strconv"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

type elf struct {
	carrying []int
	total    int
}

func parseInput(lines []string) []elf {
	var elves []elf
	current := elf{}

	for i, line := range lines {
		if len(line) == 0 {
			elves = append(elves, current)
			current = elf{}
			continue
		}

		value, _ := strconv.Atoi(line)

		current.carrying = append(current.carrying, value)
		current.total += value

		if i+1 == len(lines) {
			elves = append(elves, current)
		}
	}

	return elves
}

func solve(elves []elf) (int, int) {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].total > elves[j].total
	})

	partOne := elves[0].total
	partTwo := partOne + elves[1].total + elves[2].total

	return partOne, partTwo
}

func main() {
	path, complete := aoc.Setup(2022, 1, false)
	defer complete()

	lines := file.ToTextLines(path)
	elves := parseInput(lines)

	partOne, partTwo := solve(elves)

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}

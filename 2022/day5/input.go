package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stephensli/aoc/helpers/queue"
)

type move struct {
	from   int
	amount int
	to     int
}

func (m move) String() string {
	return fmt.Sprintf("move %d from %d to %d", m.amount, m.from, m.to)
}

func parseMoves(lines []string) []move {
	var moves []move

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")

		amount, _ := strconv.Atoi(lineSplit[1])
		from, _ := strconv.Atoi(lineSplit[3])
		to, _ := strconv.Atoi(lineSplit[5])

		moves = append(moves, move{
			from:   from,
			amount: amount,
			to:     to,
		})
	}

	return moves
}

func parseStacks(lines []string) []queue.Stack[string] {
	// determine the number of possible stacks and the index values for each
	// one of these items. Working bottom up to handle a simple parse.
	var indexPerStack []int

	for index, value := range strings.Split(lines[len(lines)-1], "") {
		if strings.TrimSpace(value) != "" {
			indexPerStack = append(indexPerStack, index)
		}
	}

	stacks := make([]queue.Stack[string], len(indexPerStack))
	for i := len(lines) - 2; i >= 0; i-- {
		line := strings.Split(lines[i], "")

		for stackIndex, index := range indexPerStack {
			if len(line) > index && strings.TrimSpace(line[index]) != "" {
				stacks[stackIndex].Push(line[index])
			}
		}
	}

	return stacks
}

func parseInput(input []string) ([]queue.Stack[string], []move) {
	stacks := parseStacks(strings.Split(input[0], "\n"))
	moves := parseMoves(strings.Split(input[1], "\n"))

	return stacks, moves
}

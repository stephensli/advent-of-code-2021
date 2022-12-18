package main

import (
	"strings"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"github.com/stephensli/aoc/helpers/queue"
)

func solution(stacks []queue.Stack[string], moves []move, batch bool) string {
	for _, move := range moves {
		batched := make([]string, move.amount, move.amount)

		for i := 0; i < move.amount; i++ {
			if !batch {
				stacks[move.to-1].Push(stacks[move.from-1].Pop())
			} else {
				batched[move.amount-1-i] = stacks[move.from-1].Pop()
			}
		}

		if batch {
			stacks[move.to-1].Push(batched...)
		}
	}

	var first string

	for _, stack := range stacks {
		first += stack.Peak()
	}

	return first
}

func main() {
	path, complete := aoc.Setup(2022, 5, false)
	defer complete()

	input := file.ToText(path)
	stacks, moves := parseInput(strings.Split(input, "\n\n"))

	first := solution(stacks, moves, false)
	aoc.PrintAnswer(1, first)

	stacks, moves = parseInput(strings.Split(input, "\n\n"))
	second := solution(stacks, moves, true)
	aoc.PrintAnswer(2, second)
}

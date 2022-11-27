package main

import "github.com/stephensli/aoc/helpers/aoc"

func main() {
	path, complete := aoc.Setup(2022, 1, true)
	defer complete()

	aoc.PrintAnswer(1, path)
}

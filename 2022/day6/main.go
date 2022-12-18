package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

// The signal is a series of seemingly-random characters that the
// device receives one at a time.
//
// Detects a start-of-packet marker in the datastream. The start of a packet is
// indicated by a sequence of four characters that are all different.
//
// The device will send your subroutine a datastream buffer (your puzzle input);
// your subroutine needs to identify the first position where the four most
// recently received characters were all different. Specifically, it needs to
// report the number of characters from the beginning of the buffer to the end
// of the first such four-character marker.
func main() {
	path, complete := aoc.Setup(2022, 6, false)
	defer complete()

	input := file.ToTextSplit(path, "")[0]

	aoc.PrintAnswer(1, solution(input, 4))
	aoc.PrintAnswer(2, solution(input, 14))
}

// First - Find the first point position in which a window of 4 characters are
// all unique. This is a travelling window.
func solution(input []string, windowSize int) int {
windowLoop:
	for i := 0; i <= len(input)-windowSize; i++ {
		window := input[i : i+windowSize]

		for i, s := range window {
			for j, s2 := range window {
				if i != j && s == s2 {
					continue windowLoop
				}
			}
		}
		return i + windowSize
	}
	return -1
}

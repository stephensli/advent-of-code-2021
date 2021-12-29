package main

import (
	"strings"

	"github.com/stephensli/advent-of-code-2021/helpers"
)

type signal struct {
	uniquePattern []string
	fourDigit     []string
}

const (
	ONE   = 2
	FOUR  = 4
	SEVEN = 3
	EIGHT = 7
)

func parseInput(lines []string) []*signal {
	var signals []*signal

	for _, line := range lines {
		splitDim := strings.Split(line, " | ")

		uniqueSignalPattern := strings.Split(splitDim[0], " ")
		fourDigit := strings.Split(splitDim[1], " ")

		signals = append(signals, &signal{
			uniquePattern: uniqueSignalPattern,
			fourDigit:     fourDigit,
		})
	}

	return signals
}

func solveEntry(sig *signal) {

}

func main() {
	// two segments are   1
	// three segments are 7
	// four segments are  4
	// seven segments are 8

	lines := helpers.ReadFileToTextLines("./day8/input-example.txt")
	signals := parseInput(lines)

	for _, val := range signals {
		solveEntry(val)
		break
	}

}

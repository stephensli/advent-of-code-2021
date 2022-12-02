package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

type Action int
type Result int
type Run []Action

var beatMap = map[Action]Action{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var typeMap = map[string]Action{
	"A": Rock,
	"B": Paper,
	"C": Scissors,

	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var typeMapPartTwo = map[string]map[string]Action{
	"X": {
		"A": Scissors,
		"B": Rock,
		"C": Paper,
	},
	"Y": {
		"A": typeMap["A"],
		"B": typeMap["B"],
		"C": typeMap["C"],
	},
	"Z": {
		"A": typeMap["B"],
		"B": typeMap["C"],
		"C": typeMap["A"],
	},
}

const (
	Rock     Action = 1
	Paper    Action = 2
	Scissors Action = 3
)

const (
	Lost Result = 0
	Draw Result = 3
	Win  Result = 6
)

func (r Run) You() Action {
	return r[1]
}

func (r Run) Them() Action {
	return r[0]
}

func (r Run) Score() int {
	return int(r.Result()) + int(r.You())
}

func (r Run) Result() Result {
	if r.You() == r.Them() {
		return Draw
	}

	if beatMap[r.You()] == r.Them() {
		return Win
	}

	return Lost
}

func parseInput(inputs [][]string) ([]Run, []Run) {
	var result []Run
	var resultPartTwo []Run

	for _, input := range inputs {
		result = append(result, Run{typeMap[input[0]], typeMap[input[1]]})
		resultPartTwo = append(resultPartTwo, Run{typeMap[input[0]], typeMapPartTwo[input[1]][input[0]]})
	}

	return result, resultPartTwo
}

func main() {
	path, complete := aoc.Setup(2022, 2, false)
	defer complete()

	lines := file.ToTextSplit(path, " ")
	runs, runsPartTwo := parseInput(lines)

	sum := 0
	sumTwo := 0

	for i := 0; i < len(runs); i++ {
		sum += runs[i].Score()
		sumTwo += runsPartTwo[i].Score()
	}

	aoc.PrintAnswer(1, sum)
	aoc.PrintAnswer(2, sumTwo)
}

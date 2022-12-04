package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"strconv"
	"strings"
)

// Every section has a unique ID number, and each Elf is assigned a range of section IDs.
//
// However, as some Elves compare their section assignments with each other, they've noticed
// that many of the assignments overlap.
//
// big list of the section assignments for each pair.
func main() {
	path, complete := aoc.Setup(2022, 4, false)
	defer complete()

	var pairs [][]assigment
	for _, ranges := range file.ToTextSplit(path, ",") {
		var pair []assigment

		for _, s := range ranges {
			rangeValues := strings.Split(s, "-")

			left, _ := strconv.Atoi(rangeValues[0])
			right, _ := strconv.Atoi(rangeValues[1])

			pair = append(pair, assigment{
				rangeLeft:  left,
				rangeRight: right,
			})
		}

		pairs = append(pairs, pair)
	}

	// part one: In how many assignment pairs does one range fully contain the other?
	// part two: In how many assignment pairs do the ranges overlap?
	var partOneCount int
	var partTwoCount int
	for _, pair := range pairs {
		if pair[0].Contains(pair[1]) || pair[1].Contains(pair[0]) {
			partOneCount += 1
		}
		if pair[0].Overlap(pair[1]) || pair[1].Overlap(pair[0]) {
			partTwoCount += 1
		}
	}

	aoc.PrintAnswer(1, partOneCount)
	aoc.PrintAnswer(2, partTwoCount)
}

package main

import (
	"fmt"
	"strings"

	"github.com/life4/genesis/slices"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"github.com/stephensli/aoc/helpers/printers"
)

func parseCaveWithAllValidMappings(inputLines [][]string) map[string]map[string]bool {
	caveMap := map[string]map[string]bool{}

	for _, entry := range inputLines {
		left := entry[0]
		right := entry[1]

		if _, ok := caveMap[left]; !ok {
			caveMap[left] = map[string]bool{}
		}

		if _, ok := caveMap[right]; !ok {
			caveMap[right] = map[string]bool{}
		}

		caveMap[left][right] = true
		caveMap[right][left] = true
	}

	return caveMap
}

func StringArrayContainsTwiceAnyLower(input []string) bool {
	if first {
		return true

	}

	countMap := map[string]int{}

	for _, s := range input {
		if s != strings.ToLower(s) {
			continue
		}

		if _, ok := countMap[s]; !ok {
			countMap[s] = 0
		}

		countMap[s] += 1

		if countMap[s] == 2 {
			return true
		}
	}

	return false
}

func findAllValidPaths(caveMap map[string]map[string]bool, nextDirection string, history []string, callback func(endHistory []string)) {
	nextPositionDirections := caveMap[nextDirection]
	history = append(history, nextDirection)

	if nextDirection == "end" {
		callback(history)
		return
	}

	for s, _ := range nextPositionDirections {
		// if we are going into a small cave and said small cave is in our history
		// then we cannot enter it again and should continue without it
		if (s == strings.ToLower(s) && slices.Contains(history, s) &&
			StringArrayContainsTwiceAnyLower(history)) || s == "start" {
			continue
		}

		newHistory := make([]string, len(history))
		copy(newHistory, history)

		findAllValidPaths(caveMap, s, history, callback)
	}
}

var first = false

func main() {
	path, deferFunc := aoc.Setup(2021, 12, false)
	defer deferFunc()

	// 1. first read the input into a format which is parseable.
	// this can be used to generate the network.
	inputLines := file.ToTextSplit(path, "-")
	caveMap := parseCaveWithAllValidMappings(inputLines)

	printers.JsonPrint(caveMap, true)

	var finalResult [][]string

	findAllValidPaths(caveMap, "start", []string{}, func(endHistory []string) {
		fmt.Println("hit end - ", endHistory)
		finalResult = append(finalResult, endHistory)
	})

	fmt.Printf("final count - %d\n", len(finalResult))
}

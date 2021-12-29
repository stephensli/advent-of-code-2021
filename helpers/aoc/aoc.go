package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Setup makes sure that we are in the current directory for calling into the aoc project.
// Simplifying the pathing. The function returned should be deferred which will result in a log of
// the task and its execution time.
func Setup(year, day int, example bool) (string, func()) {
	startTime := time.Now()

	dayName := fmt.Sprintf("day%d", day)
	path, _ := os.Getwd()

	_ = os.Chdir(filepath.Join(path, fmt.Sprintf("%d", year), dayName))
	inputFileName := fmt.Sprintf("%s.txt", dayName)

	if example {
		printExampleFlag()
		inputFileName = fmt.Sprintf("%s.example.txt", dayName)
	}

	inputPath := filepath.Join(path, strconv.Itoa(year), "inputs", inputFileName)
	fmt.Println()

	return inputPath, func() {
		fmt.Printf("AOC (%d:%d): %v\n", year, day, time.Since(startTime))

		if example {
			printExampleFlag()
		}
	}
}

// PrintAnswer  will simply log the answer in a consistent format.
func PrintAnswer(part int, answer interface{}) {
	fmt.Printf("Part %d: %v\n", part, answer)
}

func printExampleFlag() {
	fmt.Print("\n########### ###########\n# EXAMPLE # # EXAMPLE #\n########### ###########\n")
}

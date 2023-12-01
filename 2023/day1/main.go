package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"strconv"
	"strings"
)

var stringNumbers = map[string]int64{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var stringValues = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func equalOrSubStringOfValues(input string) bool {
	for _, value := range stringValues {
		if strings.HasPrefix(value, input) {
			return true
		}
	}

	return false
}

func parseNumberInput(character string) (int64, bool) {
	if value, err := strconv.ParseInt(character, 10, 64); err == nil {
		return value, true
	}

	return 0, false
}

func checkAppendNumber(partOne, partTwo []int64, char string) ([]int64, []int64, bool) {
	if value, ok := parseNumberInput(char); ok {
		if len(partOne) == 0 {
			partOne = append(partOne, value, value)
		} else {
			partOne[1] = value
		}

		if len(partTwo) == 0 {
			partTwo = append(partTwo, value, value)
		} else {
			partTwo[1] = value
		}

		return partOne, partTwo, true
	}

	return partOne, partTwo, false
}

func checkAppendDigits(array []int64, letterChain string) ([]int64, bool) {
	if !equalOrSubStringOfValues(letterChain) {
		return array, true
	}

	if value, ok := stringNumbers[strings.ToLower(letterChain)]; ok {
		if len(array) == 0 {
			array = append(array, value, value)
			return array, true
		}

		array[1] = value
		return array, true
	}

	return array, false
}

func main() {
	path, complete := aoc.Setup(2023, 1, false)
	defer complete()

	lines := file.ToTextLines(path)

	var partOneSum, partTwoSum int64

	for _, line := range lines {
		partOneNumbers := make([]int64, 0)
		partTwoNumbers := make([]int64, 0)

		var letterChain string
		var ok bool

		for _, char := range strings.Split(line, "") {
			if partOneNumbers, partTwoNumbers, ok = checkAppendNumber(partOneNumbers, partTwoNumbers, char); ok {
				letterChain = ""
				continue
			}

			letterChain += char
			partTwoNumbers, ok = checkAppendDigits(partTwoNumbers, letterChain)

			for ok {
				letterChain = letterChain[1:]
				partTwoNumbers, ok = checkAppendDigits(partTwoNumbers, letterChain)
			}
		}

		partOneSum += partOneNumbers[0]*10 + partOneNumbers[1]
		partTwoSum += partTwoNumbers[0]*10 + partTwoNumbers[1]
	}

	aoc.PrintAnswer(1, partOneSum)
	aoc.PrintAnswer(2, partTwoSum)
}

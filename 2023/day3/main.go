package main

import (
	"strconv"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

var numbers = map[string]string{
	"1": "1", "2": "2",
	"3": "3", "4": "4",
	"5": "5", "6": "6",
	"7": "7", "8": "8",
	"9": "9", "0": "0",
}

func locateConnectingNumber(lines [][]string, row, column int) string {
	value, ok := numbers[lines[row][column]]
	if !ok {
		return ""
	}

	// go left
	for i := column - 1; i >= 0; i-- {
		if leftValue, leftOk := numbers[lines[row][i]]; leftOk {
			value = leftValue + value
			continue
		}
		break
	}

	// go right
	for i := column + 1; i < len(lines); i++ {
		if rightValue, rightOk := numbers[lines[row][i]]; rightOk {
			value = value + rightValue
			continue
		}
		break
	}

	return value
}

func locateConnectingNumbers(lines [][]string, row, column int) (numbers []int64) {
	directions := [][]int{
		{row - 1, column},
		{row + 1, column},
		{row, column - 1},
		{row, column + 1},
		{row - 1, column - 1},
		{row + 1, column + 1},
		{row - 1, column + 1},
		{row + 1, column - 1},
	}

	uniquePartNumbers := map[int64]bool{}
	for _, direction := range directions {
		r, c := direction[0], direction[1]

		if c < 0 || r < 0 || r > len(lines) || c > len(lines[c]) {
			continue
		}

		if number := locateConnectingNumber(lines, r, c); number != "" {
			value, _ := strconv.ParseInt(number, 10, 64)
			uniquePartNumbers[value] = true
		}
	}

	for b := range uniquePartNumbers {
		numbers = append(numbers, b)
	}

	return numbers
}

func parser(lines [][]string) (partOne int64, partTwo int64) {
	for row, line := range lines {
		for column, v := range line {
			if v == "." {
				continue
			}

			if _, ok := numbers[v]; ok {
				continue
			}

			nums := locateConnectingNumbers(lines, row, column)
			for _, num := range nums {
				partOne += num
			}

			if v == "*" && len(nums) == 2 {
				partTwo += nums[0] * nums[1]
			}
		}
	}
	return
}

func main() {
	path, complete := aoc.Setup(2023, 3, false)
	defer complete()

	partOne, partTwo := parser(file.ToTextSplit(path, ""))

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}

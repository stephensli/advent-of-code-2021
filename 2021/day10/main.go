package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers"
	"sort"
	"strings"
)

func isOpenCharacter(input string) bool {
	switch input {
	case "(":
		fallthrough
	case "[":
		fallthrough
	case "{":
		fallthrough
	case "<":
		return true
	}

	return false
}

func inverseCharacter(input string) string {
	switch input {
	case "(":
		return ")"
	case ")":
		return "("
	case "[":
		return "]"
	case "]":
		return "["
	case "{":
		return "}"
	case "}":
		return "{"
	case "<":
		return ">"
	case ">":
		return "<"
	default:
		return "INVALID"
	}

}

func main() {
	lines := helpers.ReadFileToTextLines("./day10/input.txt")

	// count total open characters mapping
	corruptCharMap := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	notCompleteCharMap := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}

	firstSum := 0
	secondSumValues := []int{}

	for _, line := range lines {
		brackets := []string{}

		if !isOpenCharacter(string(line[0])) {
			fmt.Println("first character invalid")
			break
		}

		corrupt := false
		secondSum := 0

		for _, character := range strings.Split(line, "") {
			if isOpenCharacter(character) {
				brackets = append(brackets, character)
				continue
			}

			// if we are a closing character and our last character is not an
			// opening one for the selected character then the line is corrupted.
			if brackets[len(brackets)-1] != inverseCharacter(character) {
				fmt.Printf("line - %s\t -  Expected %s, but found %s instead.\n",
					line, inverseCharacter(brackets[len(brackets)-1]), character)

				firstSum += corruptCharMap[character]
				corrupt = true
				break
			} else {
				brackets = brackets[:len(brackets)-1]

			}
		}

		if corrupt {
			continue
		}

		// complete it

		for {
			if len(brackets) == 0 {
				break
			}

			last := brackets[len(brackets)-1]

			secondSum *= 5
			secondSum += notCompleteCharMap[inverseCharacter(last)]

			brackets = brackets[:len(brackets)-1]
		}

		secondSumValues = append(secondSumValues, secondSum)
	}

	sort.Slice(secondSumValues, func(i, j int) bool {
		return secondSumValues[i] > secondSumValues[j]
	})

	fmt.Println("first", firstSum)
	fmt.Println("second", secondSumValues[len(secondSumValues)/2])
}

package main

import (
	"github.com/life4/genesis/slices"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/cache"
	"github.com/stephensli/aoc/helpers/file"
)

var (
	Vowels = []string{"a", "e", "i", "o", "u"}
	Banned = []string{"ab", "cd", "pq", "xy"}
)

func isNicePartTwo(input []string) bool {
	twoLetterCache := cache.New[string, int]()
	endIndexCache := cache.New[int, bool]()
	repeatWithSpaceCount := 0

	for i := 0; i < len(input); i++ {
		if i == len(input)-1 {
			continue
		}

		current := input[i]
		next := input[i+1]

		// It contains at least one letter which repeats with exactly one letter
		// between them, like xyx, abcdefeghi (efe), or even aaa.
		//
		// do this before we do a skip action on the following.
		if i > 0 {
			previous := input[i-1]

			if next == previous {
				repeatWithSpaceCount += 1
			}
		}

		// It contains a pair of any two letters that appears at least twice in the
		// string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not
		// like aaa (aa, but it overlaps).
		if current == next {
			// if the previous one was the end of a previous group, and
			// we match the previous value continue. Allowing us to skip
			// groups of three being classed as two valid groups.
			if i > 0 && endIndexCache.Has(i-1) {
				continue
			}

			endIndexCache.Set(i, true)
		}

		key := current + next
		val, _ := twoLetterCache.GetOrSet(key, 0)
		twoLetterCache.Set(key, val+1)

	}

	anyTwoLettersAppearingTwice := slices.Any(twoLetterCache.Values(), func(elm int) bool {
		return elm >= 2
	})

	return anyTwoLettersAppearingTwice && repeatWithSpaceCount >= 1
}

func isNice(input []string) bool {
	vowelCount := 0
	twiceCount := 0

	noInvalidStrings := true

	for i := 0; i < len(input); i++ {
		if slices.Contains(Vowels, input[i]) {
			vowelCount += 1
		}

		if i != len(input)-1 {
			current := input[i]
			next := input[i+1]

			if current == next {
				twiceCount += 1
				continue
			}

			if slices.Contains(Banned, current+next) {
				noInvalidStrings = false
				return false
			}

		}

	}

	return vowelCount >= 3 && twiceCount >= 1 && noInvalidStrings
}

func solution(input [][]string, first bool) int {
	return slices.Reduce(input, 0, func(elm []string, acc int) int {
		if first && isNice(elm) {
			acc += 1
		} else if isNicePartTwo(elm) {
			acc += 1
		}

		return acc
	})
}

func main() {
	path, deferFunc := aoc.Setup(2015, 5, false)
	defer deferFunc()

	input := file.ToTextSplit(path, "")

	aoc.PrintAnswer(1, solution(input, true))
	aoc.PrintAnswer(2, solution(input, false))
}

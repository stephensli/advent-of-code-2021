package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/stephensli/advent-of-code-2021/helpers"
)

func parseInput(input []string) (string, map[string]string) {
	mappings := map[string]string{}

	return input[0], mappings
}

func first(elements map[string]int64) {
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	sum := int64(0)

	for _, v := range elements {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}

	fmt.Printf("result: %d - %d = %d, length = %d\n",
		max, min, max-min, sum)
}

func main() {
	lines := helpers.ReadFileToTextLines("./day14/input-example.txt")
	translate := map[string]string{}
	polymer := lines[0]

	for _, v := range lines[2:] {
		split := strings.Split(v, " -> ")
		translate[split[0]] = split[1]
	}

	pairs := map[string]int64{}
	elements := map[string]int64{}

	for i := 0; i < len(polymer)-1; i++ {
		p := polymer[i : i+2]

		helpers.SetIfMissing(pairs, p, 0)
		pairs[p] += 1
	}

	for k := 0; k < len(polymer); k++ {
		p := string(polymer[k])

		helpers.SetIfMissing(elements, p, 0)
		elements[p] += 1
	}

	for i := 0; i < 40; i++ {
		newPairs := map[string]int64{}

		for pairKey := range pairs {
			insert := translate[pairKey]
			c := pairs[pairKey]

			helpers.SetIfMissing(newPairs, string(pairKey[0])+insert, 0)
			newPairs[string(pairKey[0])+insert] += c

			helpers.SetIfMissing(newPairs, insert+string(pairKey[1]), 0)
			newPairs[insert+string(pairKey[1])] += c

			helpers.SetIfMissing(elements, insert, 0)
			elements[insert] += c
		}

		pairs = newPairs
	}

	first(elements)
}

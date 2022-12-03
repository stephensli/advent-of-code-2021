package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers/cache"
	"strconv"
	"unicode"
)

type item rune

func (i item) String() string {
	return fmt.Sprintf("%v: %s", string(i), strconv.Itoa(i.priority()))
}

// priority helps prioritize item rearrangement, every item type can be converted to a priority:
//
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func (i item) priority() int {
	if unicode.IsUpper(rune(i)) {
		return int(i) - 64 + 26
	}

	return int(i) - 96
}

type rucksack struct {
	compartments [][]item
	existsInBoth map[item]int
}

func newRucksack(items string) *rucksack {
	var compartment []item
	var compartments [][]item
	existsInBoth := map[item]int{}

	historyMap := cache.New[rune, bool]()

	for i, r := range []rune(items) {
		compartment = append(compartment, item(r))

		if i+1 <= len(items)/2 {
			historyMap.Set(r, true)
		} else if historyMap.Has(r) {
			if _, ok := existsInBoth[item(r)]; ok {
				existsInBoth[item(r)] += 1
			} else {
				existsInBoth[item(r)] = 1
			}
		}

		if i+1 == len(items)/2 || i+1 == len(items) {
			compartments = append(compartments, compartment)
			compartment = []item{}
		}

	}

	return &rucksack{
		compartments: compartments,
		existsInBoth: existsInBoth,
	}
}

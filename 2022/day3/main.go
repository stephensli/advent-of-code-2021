package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

// Each rucksack has two large compartments. All items of a given type are meant to go into exactly
// one of the two compartments.
//
// Every item type is identified by a single lowercase or uppercase letter (that is, `a` and `A`
// refer to different types of items).
//
// The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.
func main() {
	path, complete := aoc.Setup(2022, 3, false)
	defer complete()

	var rucksacks []*rucksack
	for _, s := range file.ToTextLines(path) {
		rucksacks = append(rucksacks, newRucksack(s))
	}

	// Part one: Find the total sum of all types that exist in both compartments. This is determined
	// on creation of the rucksack and thus only requires now adding all the values together.
	var partOnePriorityTotal int

	for _, r := range rucksacks {
		for val := range r.existsInBoth {
			partOnePriorityTotal += val.priority()
		}
	}

	// Part two: Find the item that exists in all batches of three rucksacks, only one item will
	// exist in all three batched rucksacks. Sum the total of these priorities together to work
	// out the value.
	var partTwoPriorityTotal int

	// Step in batches of three, scopes down the required amount of work to what we already know
	// is our data set. Iterating each selection to find what exists in all three. (map of count
	// per rucksack.
	for i := 2; i < len(rucksacks); i += 3 {
		rucksackSet := []*rucksack{rucksacks[i-2], rucksacks[i-1], rucksacks[i]}
		count := map[item]int{}

		for i, r := range rucksackSet {
			// if we have already seen it when looping this sack then we don't want to include it
			// within the following checks. Otherwise, it will result in higher than 3x total
			// counts.
			thisSackSeen := map[item]bool{}

			for _, compartment := range r.compartments {
				for _, item := range compartment {
					if thisSackSeen[item] {
						continue
					}

					thisSackSeen[item] = true

					if _, ok := count[item]; ok {
						count[item] += 1
					} else {
						count[item] = 1
					}

					if i == 2 && count[item] == 3 {
						partTwoPriorityTotal += item.priority()
					}
				}
			}
		}
	}

	aoc.PrintAnswer(1, partOnePriorityTotal)
	aoc.PrintAnswer(2, partTwoPriorityTotal)
}

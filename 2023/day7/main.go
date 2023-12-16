package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

type handType string

var (
	fiveOfAKindType  handType = "FIVE_OF_A_KIND"
	fourOfAKindType  handType = "FOUR_OF_A_KIND"
	fullHouseType    handType = "FULL_HOUSE"
	threeOfAKindType handType = "THREE_OF_A_KIND"
	twoPairType      handType = "TWO_PAIR"
	onePairType      handType = "ONE_PAIR"
	highCardType     handType = "HIGH_CARD"
)

var typeStrengthRanking = map[handType]int{
	fiveOfAKindType:  1,
	fourOfAKindType:  2,
	fullHouseType:    3,
	threeOfAKindType: 4,
	twoPairType:      5,
	onePairType:      6,
	highCardType:     7,
}

var strengthMapping = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
}

type cardHand struct {
	cardType handType
	values   []string
	value    int64
}

// determineCardType sets the card type on the card hand and also returns
// the value. Attempts to execute the determination in a single loop.
func (c *cardHand) determineCardType(p2 bool) {
	kindCountMap := map[string]int{}
	var uniqueValues []string

	for _, v := range c.values {
		if _, ok := kindCountMap[v]; !ok {
			uniqueValues = append(uniqueValues, v)
		}
		kindCountMap[v] += 1
	}

	// Now lets go and pull the value for J, set the value for J to zero and
	// then perform this sort. This will place the next highest and meaningful
	// value to the top which will now accept the J value.
	var jokerCardCount int

	if p2 && kindCountMap["J"] != 5 {
		jokerCardCount = kindCountMap["J"]
		kindCountMap["J"] = 0
	}

	sort.Slice(uniqueValues, func(i, j int) bool {
		return kindCountMap[uniqueValues[i]] > kindCountMap[uniqueValues[j]]
	})

	// Now go and apply the joker value to the next highest.
	if p2 && jokerCardCount != 0 {
		kindCountMap[uniqueValues[0]] += jokerCardCount
	}

	// If the highest ranked value is 1, then we know all values are unique.
	if kindCountMap[uniqueValues[0]] == 1 {
		c.cardType = highCardType
		return
	}

	// If the highest ranked value is 5, then we know we have five of a kind.
	if kindCountMap[uniqueValues[0]] == 5 {
		c.cardType = fiveOfAKindType
		return
	}

	// If the highest ranked value is 4, then we know we have four of a kind.
	if kindCountMap[uniqueValues[0]] == 4 {
		c.cardType = fourOfAKindType
		return
	}

	// If the highest ranked value is 3, and the second is 2, then we have a full
	// house
	if kindCountMap[uniqueValues[0]] == 3 && kindCountMap[uniqueValues[1]] == 2 {
		c.cardType = fullHouseType
		return
	}

	// If the highest ranked value is 3, and the second is 2, then we have a three
	// of a kind.
	if kindCountMap[uniqueValues[0]] == 3 && kindCountMap[uniqueValues[1]] == 1 {
		c.cardType = threeOfAKindType
		return
	}

	// If the highest is two, and the second highest is two, we have two pair.
	if kindCountMap[uniqueValues[0]] == 2 && kindCountMap[uniqueValues[1]] == 2 {
		c.cardType = twoPairType
		return
	}

	// high card count is already determined, so this must be one pair.
	c.cardType = onePairType
	return
}

func parser(input []string, p2 bool) (hands []cardHand) {
	for _, v := range input {
		hand := cardHand{}

		for i2, v2 := range strings.Split(v, " ") {
			if i2 == 0 {
				hand.values = strings.Split(v2, "")
				continue
			}

			value, _ := strconv.ParseInt(v2, 10, 64)
			hand.value = value
		}

		hand.determineCardType(p2)
		hands = append(hands, hand)
	}

	// sort the decks by strengh
	sort.Slice(hands, func(i, j int) bool {
		for ii := range hands[i].values {
			// first sort by card type and then sort by card values
			if hands[i].cardType != hands[j].cardType {
				return typeStrengthRanking[hands[i].cardType] > typeStrengthRanking[hands[j].cardType]
			}

			firstCardStrength := strengthMapping[hands[i].values[ii]]
			secondCardStrength := strengthMapping[hands[j].values[ii]]

			if firstCardStrength == secondCardStrength {
				continue
			}

			return firstCardStrength > secondCardStrength
		}

		return true
	})

	return hands
}

// Your goal is to order them based on the strength of each hand.
// 5 cards per hand
//
// - Five of a kind, where all five cards have the same label: AAAAA
//
// - Four of a kind, where four cards have the same label and one card has a different label: AA8AA
//
// - Full house, where three cards have the same label, and the remaining two
// cards share a different label: 23332
//
// - Three of a kind, where three cards have the same label, and the remaining two
// cards are each different from any other card in the hand: TTT98
//
// - Two pair, where two cards share one label, two other cards share a second
// label, and the remaining card has a third label: 23432
//
// - One pair, where two cards share one label, and the other three cards have a
// different label from the pair and each other: A23A4
//
// - High card, where all cards' labels are distinct: 23456
func main() {
	path, complete := aoc.Setup(2023, 7, false)
	defer complete()

	var partOne int64
	var partTwo int64

	for rank, ch := range parser(file.ToTextLines(path), false) {
		partOne += (int64(rank) + 1) * ch.value
	}

	// Part Two
	strengthMapping["J"] = 14
	for rank, ch := range parser(file.ToTextLines(path), true) {
		partTwo += (int64(rank) + 1) * ch.value
	}

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}

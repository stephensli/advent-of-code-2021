package main

import (
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"strconv"
	"strings"
)

type Card struct {
	id               int64
	winningNumbers   []int64
	scratchedNumbers []int64
	matchingNumbers  []int64
	partOneScore     int
	copies           int
}

func parse(lines []string) (cards []*Card) {
	for _, line := range lines {
		cardAndScratchNumbers := strings.Split(line, " | ")

		cardSplit := strings.Split(cardAndScratchNumbers[0], ": ")
		id, _ := strconv.ParseInt(strings.TrimPrefix(strings.ToLower(cardSplit[0]), "card "), 10, 64)

		winningNumbers := make([]int64, 0)
		scratchNumbers := make([]int64, 0)
		matchingNumbers := make([]int64, 0)
		partOneScore := 0

		for _, s := range strings.Split(strings.TrimSpace(cardSplit[1]), " ") {
			if s == "" {
				continue
			}

			number, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
			winningNumbers = append(winningNumbers, number)
		}

		for _, s := range strings.Split(strings.TrimSpace(cardAndScratchNumbers[1]), " ") {
			if s == "" {
				continue
			}

			number, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
			scratchNumbers = append(scratchNumbers, number)
		}

		for _, number := range winningNumbers {
			for _, targetNumber := range scratchNumbers {
				if number == targetNumber {
					matchingNumbers = append(matchingNumbers, number)

					if partOneScore == 0 {
						partOneScore += 1
						break
					}

					partOneScore += partOneScore
				}
			}
		}

		cards = append(cards, &Card{
			id:               id,
			winningNumbers:   winningNumbers,
			scratchedNumbers: scratchNumbers,
			matchingNumbers:  matchingNumbers,
			partOneScore:     partOneScore,
			copies:           1,
		})
	}

	return cards
}

func main() {
	path, complete := aoc.Setup(2023, 4, false)
	defer complete()

	cards := parse(file.ToTextLines(path))

	p1 := 0
	p2 := 0

	for i, card := range cards {
		p1 += card.partOneScore
		p2 += card.copies

		for j := 0; j < len(card.matchingNumbers); j++ {
			cards[i+1+j].copies += card.copies
		}
	}

	aoc.PrintAnswer(1, p1)
	aoc.PrintAnswer(2, p2)
}

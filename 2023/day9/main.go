package main

import (
	"strconv"
	"strings"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

type Record struct {
	history [][]int64
}

func (h *Record) computePreviousHistoryValues() {
	h.history[len(h.history)-1] = append([]int64{0}, h.history[len(h.history)-1]...)

	for i := len(h.history) - 2; i >= 0; i-- {
		currentValue := h.history[i][0]
		previousValue := h.history[i+1][0]

		h.history[i] = append([]int64{currentValue - previousValue}, h.history[i]...)
	}
}

func (h *Record) computeNextHistoryValues() {
	h.history[len(h.history)-1] = append(h.history[len(h.history)-1], 0)

	for i := len(h.history) - 2; i >= 0; i-- {
		currentValue := h.history[i][len(h.history[i])-1]
		previousValue := h.history[i+1][len(h.history[i+1])-1]

		h.history[i] = append(h.history[i], currentValue+previousValue)
	}
}

func (h *Record) computeHistory() {
	allZeroValues := false
	current := h.history[0]

	for !allZeroValues {
		next := []int64{}
		totalZero := 0

		for i := 0; i < len(current)-1; i++ {
			nextValue := current[i+1] - current[i]
			next = append(next, nextValue)

			if nextValue == 0 {
				totalZero += 1
			}
		}

		h.history = append(h.history, next)
		current = next

		if totalZero == len(next) {
			break
		}
	}
}

func parser(input []string) (records []*Record) {
	for _, v := range input {
		var base []int64

		for _, value := range strings.Split(v, " ") {
			intValue, _ := strconv.ParseInt(value, 10, 64)
			base = append(base, intValue)
		}

		record := Record{history: [][]int64{base}}
		record.computeHistory()

		records = append(records, &record)
	}

	return records
}

func main() {
	path, complete := aoc.Setup(2023, 9, false)
	defer complete()

	var partOne int64
	var partTwo int64

	for _, r := range parser(file.ToTextLines(path)) {
		r.computeNextHistoryValues()
		r.computePreviousHistoryValues()

		partOne += r.history[0][len(r.history[0])-1]
		partTwo += r.history[0][0]
	}

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}


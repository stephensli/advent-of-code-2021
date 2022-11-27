package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers"
	"strconv"
	"strings"
	"time"
)

type bitValues struct {
	zero, one int
}

func (b bitValues) getMostCommonBit() string {
	if b.zero > b.one {
		return "0"
	}
	return "1"
}

func (b bitValues) getLeastCommonBit() string {
	if b.zero > b.one {
		return "1"
	}
	return "0"
}

func generateBitsFromInput(input []string) []*bitValues {
	bits := make([]*bitValues, len(input[0]))

	for _, str := range input {
		for idx, bit := range strings.Split(str, "") {
			if bits[idx] == nil {
				bits[idx] = &bitValues{0, 0}
			}

			if bit == "0" {
				bits[idx].zero += 1
			} else {
				bits[idx].one += 1
			}
		}
	}

	return bits
}

func first(input []string) int64 {
	powerConsumption := int64(0)

	bits := generateBitsFromInput(input)

	gammaBinary := ""
	epsilonBinary := ""

	for _, bit := range bits {
		gammaBinary += bit.getMostCommonBit()
		epsilonBinary += bit.getLeastCommonBit()
	}

	parsedGamma, _ := strconv.ParseInt(gammaBinary, 2, 64)
	parsedEpsilon, _ := strconv.ParseInt(epsilonBinary, 2, 64)
	powerConsumption += parsedGamma * parsedEpsilon

	return powerConsumption
}

func second(input []string) int64 {
	// 1. oxygen generator rating.
	// determine the most common value for that bit position.
	// keep only values that have that bit in that position.
	// if both 1 and 0 are equal, keep 1.
	index := 0

	start := time.Now()

	oxygenInput := input

	for {
		if len(oxygenInput) == 1 {
			break
		}

		bits := generateBitsFromInput(oxygenInput)

		mostCommon := bits[index].getMostCommonBit()
		newInput := []string{}

		// filter out input for all most common
		for _, val := range oxygenInput {
			if string(val[index]) == mostCommon {
				newInput = append(newInput, val)
			}
		}

		oxygenInput = newInput
		index += 1
	}

	// 2. co2 scrubber rating
	// keep the numbers of which have the least common
	scrubberInput := input
	index = 0

	for {
		if len(scrubberInput) == 1 {
			break
		}

		bits := generateBitsFromInput(scrubberInput)

		leastCommon := bits[index].getLeastCommonBit()
		newInput := []string{}

		// filter out input for all most common
		for _, val := range scrubberInput {
			if string(val[index]) == leastCommon {
				newInput = append(newInput, val)
			}
		}

		scrubberInput = newInput
		index += 1
	}

	oxygenInputValue, _ := strconv.ParseInt(oxygenInput[0], 2, 64)
	scrubberInputValue, _ := strconv.ParseInt(scrubberInput[0], 2, 64)

	fmt.Println("time: ", time.Since(start))

	return oxygenInputValue * scrubberInputValue
}

func main() {
	input := helpers.ReadFileToTextLines("./day3/input.txt")

	fmt.Println("first", first(input))
	fmt.Println("second", second(input))
}

package main

import (
	"github.com/life4/genesis/slices"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"strconv"
	"strings"
)

type convertMappingValue struct {
	destinationStartRange int64
	sourceStartRange      int64
	rangeLength           int64
}

type convertType string

var (
	seedToSoil            convertType = "seed-to-soil"
	soilToFertilizer      convertType = "soil-to-fertilizer"
	fertilizerToWater     convertType = "fertilizer-to-water"
	waterTolight          convertType = "water-to-light"
	lightToTemperature    convertType = "light-to-temperature"
	temperatureToHumidity convertType = "temperature-to-humidity"
	humidityToLocation    convertType = "humidity-to-location"
)

var transferMap = []convertType{
	seedToSoil,
	soilToFertilizer,
	fertilizerToWater,
	waterTolight,
	lightToTemperature,
	temperatureToHumidity,
	humidityToLocation,
}

type convertMappings map[convertType][]convertMappingValue

func (c convertMappings) getMappingNumberForSeed(seed int64, mappingType convertType) int64 {
	for _, value := range c[mappingType] {
		if seed < value.sourceStartRange {
			continue
		}

		if seed > value.sourceStartRange+value.rangeLength-1 {
			continue
		}

		return value.destinationStartRange + (seed - value.sourceStartRange)
	}

	return seed
}

func (c convertMappings) getMappingNumberForReverseSeed(seed int64, mappingType convertType) int64 {
	for _, value := range c[mappingType] {
		if seed < value.destinationStartRange {
			continue
		}

		if seed > value.destinationStartRange+value.rangeLength-1 {
			continue
		}

		return value.sourceStartRange + (seed - value.destinationStartRange)
	}

	return seed
}

func parser(lines []string) (seeds []int64, mappings convertMappings) {
	mappings = convertMappings{}

	for _, s := range strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ") {
		value, _ := strconv.ParseInt(s, 10, 64)
		seeds = append(seeds, value)
	}

	resetMapping := true
	mappingIndex := convertType("")

	for _, line := range lines[2:] {
		if resetMapping {
			resetMapping = false
			mappingIndex = convertType(strings.Split(line, " ")[0])
			mappings[mappingIndex] = make([]convertMappingValue, 0)
			continue
		}

		if line == "" {
			resetMapping = true
			continue
		}

		mapping := convertMappingValue{}

		for i, s := range strings.Split(line, " ") {
			if s == " " {
				continue
			}

			value, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)

			switch i {
			case 0:
				mapping.destinationStartRange = value
			case 1:
				mapping.sourceStartRange = value
			case 2:
				mapping.rangeLength = value
			}
		}

		mappings[mappingIndex] = append(mappings[mappingIndex], mapping)
	}

	return seeds, mappings
}

func main() {
	path, complete := aoc.Setup(2023, 5, false)
	defer complete()

	seeds, mapping := parser(file.ToTextLines(path))

	var partOne int64
	var partTwo int64

	for i, seed := range seeds {
		currentValue := seed
		for _, c := range transferMap {
			currentValue = mapping.getMappingNumberForSeed(currentValue, c)
		}

		if i == 0 || currentValue < partOne {
			partOne = currentValue
		}
	}

	var seedSets [][]int64
	var seedSet []int64

	for _, seed := range seeds {
		seedSet = append(seedSet, seed)

		if len(seedSet) >= 2 {
			seedSets = append(seedSets, seedSet)
			seedSet = []int64{}
		}
	}

	var i int64
	for partTwo == 0 {
		i += 1

		currentValue := i

		for _, c := range slices.Reverse(transferMap) {
			currentValue = mapping.getMappingNumberForReverseSeed(currentValue, c)
		}

		for _, set := range seedSets {
			if currentValue >= set[0] && currentValue <= set[0]+set[1] {
				partTwo = i
				break
			}
		}

		if partTwo != 0 {
			break
		}
	}

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}

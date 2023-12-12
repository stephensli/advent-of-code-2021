package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

type Race struct {
	time     int64
	distance int64
}

func parser(input []string) (partOne []Race, partTwo Race) {
	time := strings.Split(input[0], " ")[1:]
	distance := strings.Split(input[1], " ")[1:]

	partTwoTime, _ := strconv.ParseInt(strings.Join(time, ""), 10, 64)
	partTwoDistance, _ := strconv.ParseInt(strings.Join(distance, ""), 10, 64)

	partTwo = Race{
		time:     partTwoTime,
		distance: partTwoDistance,
	}

	for i := 0; i < len(time); i++ {
		t, _ := strconv.ParseInt(time[i], 10, 64)
		d, _ := strconv.ParseInt(distance[i], 10, 64)

		partOne = append(partOne, Race{
			time:     t,
			distance: d,
		})
	}
	return
}

func findValueForRace(r Race) int64 {
	var largest, smallest int64
	wg := sync.WaitGroup{}

	wg.Add(2)

	// find the largest value that can pass the distance
	go func() {
		defer wg.Done()
		for i := r.time; i > 0; i-- {
			difference := r.time - i
			distance := i * difference
			if distance > r.distance {
				largest = i
				break
			}
		}
	}()

	// find the first value that can pass the distance
	go func() {
		defer wg.Done()
		for j := int64(0); j < r.time; j++ {
			difference := r.time - j
			distance := j * difference
			if distance > r.distance {
				smallest = j
				break
			}
		}
	}()

	wg.Wait()
	return largest - smallest + 1
}

func main() {
	path, complete := aoc.Setup(2023, 6, false)
	defer complete()

	partOneRaces, partTwoRace := parser(file.ToTextLines(path))
	var partOne int64

	for i, r := range partOneRaces {
		value := findValueForRace(r)

		if i == 0 {
			partOne = value
		} else {
			partOne *= value
		}
	}

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, findValueForRace(partTwoRace))
}

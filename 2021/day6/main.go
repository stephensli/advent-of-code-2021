package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers"
	"strconv"
	"strings"
)

type Fish struct {
	value, delta int
}

func processDays(fish map[int]*Fish) (newFish bool) {
	for i := 8; i > 0; i-- {
		fish[i-1].delta += fish[i].value
	}

	for i := 8; i >= 0; i-- {
		if i == 0 {
			if fish[i].delta != 0 || fish[i].value != 0 {
				fish[8].value += fish[i].value
				fish[6].value += fish[i].value
				fish[i].value = fish[i].delta
				fish[i].delta = 0
			}
			continue
		}

		fish[i].value = fish[i].delta
		fish[i].delta = 0
	}

	//	printFishes(fish)
	return false
}

func printFishes(fishes map[int]*Fish) {
	fmt.Println("----")
	for i, fish := range fishes {
		fmt.Println(i, fish)
	}
}

func main() {
	input := helpers.ReadFileToTextLines("./day6/input.txt")

	fish := map[int]*Fish{
		0: {value: 0, delta: 0},
		1: {value: 0, delta: 0},
		2: {value: 0, delta: 0},
		3: {value: 0, delta: 0},
		4: {value: 0, delta: 0},
		5: {value: 0, delta: 0},
		6: {value: 0, delta: 0},
		7: {value: 0, delta: 0},
		8: {value: 0, delta: 0},
	}

	for _, val := range strings.Split(input[0], ",") {
		num, _ := strconv.Atoi(val)
		fish[num].value += 1
	}

	fmt.Println("start")
	printFishes(fish)

	days := 256
	count := 0

	for {
		count += 1
		processDays(fish)

		if count == days {
			break
		}
	}

	sum := 0
	for _, fish := range fish {
		sum += fish.value
	}

	fmt.Println("result", sum)
}

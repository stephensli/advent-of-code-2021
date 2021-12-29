package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func first(input []int) int {
	totalIncrease := 0
	previous := 0

	first := true

	for _, num := range input {
		if first {
			first = false
			previous = num
			continue
		}

		if previous < num {
			totalIncrease += 1
		}

		previous = num
	}

	return totalIncrease
}

func second(input []int) int {
	totalIncrease := 0
	previous := 0
	first := true

	for idx, _ := range input {
		if idx+2 >= len(input) {
			break
		}

		next := input[idx] + input[idx+1] + input[idx+2]

		if first {
			previous = next
			first = false
		}

		if previous < next {
			totalIncrease += 1
		}

		previous = next
	}

	return totalIncrease
}

func main() {
	file, err := os.Open("./day1/input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []int

	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		input = append(input, num)
	}

	fmt.Println("first", first(input))
	fmt.Println("second", second(input))
}

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileInput(filePath string) ([]int, []*Board) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputValues := []int{}

	var boards []*Board

	currentIndex := 0
	currentBoard := &Board{
		content:       [5][5]int{},
		matchedValues: map[int]bool{},
	}

	first := true

	for scanner.Scan() {
		lineStr := scanner.Text()

		if first {
			if lineStr == "" {
				first = false
				continue
			}

			for _, v := range strings.Split(lineStr, ",") {
				num, _ := strconv.Atoi(v)
				inputValues = append(inputValues, num)
			}

			continue
		}

		if lineStr == "" {
			boards = append(boards, currentBoard)

			currentBoard = &Board{
				content:       [5][5]int{},
				matchedValues: map[int]bool{},
			}

			currentIndex = 0
		} else {
			offset := 0

			for i, v := range strings.Split(lineStr, " ") {
				if strings.TrimSpace(v) == "" {
					offset += 1
					continue
				}

				num, _ := strconv.Atoi(v)
				currentBoard.content[currentIndex][i-offset] = num
			}

			currentIndex += 1
		}
	}

	boards = append(boards, currentBoard)

	return inputValues, boards
}

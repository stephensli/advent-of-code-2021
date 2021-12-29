package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readText(scanner *bufio.Scanner, actionFunc func(input string)) {
	lineStr := scanner.Text()
	actionFunc(lineStr)
}

func readFile(filePath string, actionFunc func(input string)) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		readText(scanner, actionFunc)
	}
}

func ToText(filePath string) string {
	var response []string

	readFile(filePath, func(input string) {
		response = append(response, input)
	})

	return strings.Join(response, "\n")
}

func ToTextLines(filePath string) []string {
	var response []string

	readFile(filePath, func(input string) {
		response = append(response, input)
	})

	return response
}

func ToNumbers(filePath string) []int {
	var response []int

	readFile(filePath, func(input string) {
		num, _ := strconv.Atoi(input)

		response = append(response, num)
	})

	return response
}

func ToTextSplit(filePath string, split string) [][]string {
	var response [][]string

	readFile(filePath, func(input string) {
		var line []string

		for _, val := range strings.Split(input, split) {
			line = append(line, val)
		}

		response = append(response, line)
	})

	return response
}

func ToNumbersSplit(filePath string, split string) [][]int {
	var response [][]int

	readFile(filePath, func(input string) {
		line := []int{}

		for _, val := range strings.Split(input, split) {
			num, _ := strconv.Atoi(val)
			line = append(line, num)
		}

		response = append(response, line)
	})

	return response
}

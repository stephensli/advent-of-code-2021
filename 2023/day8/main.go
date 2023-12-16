package main

import (
	"fmt"
	"strings"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"github.com/stephensli/aoc/helpers/maths"
)

func parser(input []string) ([]string, []string, map[string]string) {
	code := strings.Split(input[0], "")
	lines := input[2:]

	navigationMap := map[string]string{}
	var nodesEndingWithA []string

	for _, v := range lines {
		line := strings.Split(v, " = ")

		source := line[0]
		target := strings.Split(line[1][1:len(line[1])-1], ", ")

		navigationMap[fmt.Sprintf("%s:%s", source, "L")] = target[0]
		navigationMap[fmt.Sprintf("%s:%s", source, "R")] = target[1]

		if strings.HasSuffix(source, "A") {
			nodesEndingWithA = append(nodesEndingWithA, source)
		}
	}

	return code, nodesEndingWithA, navigationMap
}

func main() {
	path, complete := aoc.Setup(2023, 8, false)
	defer complete()

	code, endingWithA, navigationMap := parser(file.ToTextLines(path))

	partOneSteps := 0
	source := "AAA"

	for source != "ZZZ" {
		codePosition := code[partOneSteps%len(code)]
		partOneSteps += 1

		if source = navigationMap[fmt.Sprintf("%s:%s", source, codePosition)]; source == "" {
			panic("could not locate source value")
		}
	}

	done := make(chan interface{})
	output := make(chan int)

	findFirstZValue := func(
		output chan int,
		source string,
		codes []string,
		navigation map[string]string) {
		partTwoSteps := 0

		for {
			codePosition := code[partTwoSteps%len(code)]
			if source = navigationMap[fmt.Sprintf("%s:%s", source, codePosition)]; source == "" {
				panic("could not locate source value")
			}

			partTwoSteps += 1
			if strings.HasSuffix(source, "Z") {
				output <- partTwoSteps
				return
			}
		}
	}

	lowestZValues := []int{}

	go func() {
		for e := range output {
			lowestZValues = append(lowestZValues, e)

			if len(lowestZValues) == len(endingWithA) {
				close(done)
				return
			}
		}
	}()

	for _, v := range endingWithA {
		go findFirstZValue(output, v, code, navigationMap)
	}

	<-done
	fmt.Printf("lowestZValues: %v\n", lowestZValues)

	aoc.PrintAnswer(1, partOneSteps)
	aoc.PrintAnswer(2, maths.LowestCommonMultiple(lowestZValues[0], lowestZValues[1], lowestZValues[1:]...))
}

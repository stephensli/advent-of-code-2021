package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
	"strings"
)

func solution(input string, leadingZeros string) int {
	number := 0

	for {
		inputString := fmt.Sprintf("%s%d", input, number)
		hash := md5.Sum([]byte(inputString))
		value := hex.EncodeToString(hash[:])

		if strings.HasPrefix(value, leadingZeros) {
			return number
		}
		number += 1
	}
	return -1
}

func main() {
	path, deferFunc := aoc.Setup(2015, 4, false)
	defer deferFunc()

	input := file.ToTextLines(path)[0]

	aoc.PrintAnswer(1, solution(input, "00000"))
	aoc.PrintAnswer(2, solution(input, "000000"))
}

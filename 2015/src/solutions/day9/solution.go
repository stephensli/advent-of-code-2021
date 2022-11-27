package main

import (
	"fmt"
	"github.com/stephensli/aoc/helpers/aoc"
)

func main() {
	path, deferFunc := aoc.Setup(2015, 6, true)
	defer deferFunc()

	fmt.Println(path)
}

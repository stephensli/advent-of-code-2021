package main

import (
	"fmt"
	"math"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

// The filesystem consists of a tree of files (plain data) and directories
// (which can contain other directories or files). The outermost directory is
// called /.
//
// Within the terminal output, lines that begin with $ are commands you
// executed, very much like some modern computers.
//
// guessed: 1,235,336, 1,539,532, 1,266,594, 1,648,397
//
// problem: was caching based on name, name is not unique and so it's now based
// on the base64 encoding of the parent name(s) and the current name.
func main() {
	path, complete := aoc.Setup(2022, 7, false)
	defer complete()

	root := parseInput(file.ToTextLines(path))
	directories := directorySizes(root)

	requiredSpace := int64(30_000_000)
	fileSystem := int64(70_000_000)

	outmostDir := directories[0].size
	deleteTarget := requiredSpace - (fileSystem - outmostDir)

	var first int64
	for _, val := range directories {
		if val.size <= 100_000 {
			first += val.size
		}
	}

	second := int64(math.MaxInt64)
	for _, directory := range directories {
		if directory.size >= deleteTarget && directory.size < second {
			fmt.Println("directory=", directory)
			second = directory.size
		}
	}

	aoc.PrintAnswer(1, first)
	aoc.PrintAnswer(2, second)
}

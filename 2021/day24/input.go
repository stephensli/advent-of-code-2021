package main

import (
	"github.com/stephensli/aoc/helpers/file"
)

func parseInput(filePath string) [][]string {
	lines := file.ToTextLines(filePath)

	var sections [][]string

	section := []string{lines[0]}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		if line[0] == 'i' {
			sections = append(sections, section)
			section = []string{line}
		} else {
			section = append(section, line)
		}
	}
	sections = append(sections, section)
	// printers.JsonPrint(sections, true)
	return sections
}

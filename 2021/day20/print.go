package main

import "fmt"

func printInfiniteGrid(input map[Position]int, min, max Position) {

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if val, ok := input[Position{x, y}]; ok && val == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

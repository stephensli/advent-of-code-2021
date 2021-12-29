package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers"
	"strconv"
	"strings"
)

// Position - The position in the infinite image plane.
type Position struct {
	x, y int
}

func parseInputImage(inputLines []string) [][]int {
	output := make([][]int, len(inputLines))

	for i, line := range inputLines {
		output[i] = make([]int, len(line))

		for j, value := range strings.Split(line, "") {
			outputVal := 0

			if value == "#" {
				outputVal = 1
			}

			output[i][j] = outputVal
		}
	}

	return output
}

func getPixelValueFromGrid(inputGrid map[Position]int, targetPosition Position) int {
	combined := ""

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			y := targetPosition.y - 1 + i
			x := targetPosition.x - 1 + j

			if val, ok := inputGrid[Position{x, y}]; ok && val == 1 {
				combined += "1"
			} else {
				combined += "0"
			}
		}
	}

	val, _ := strconv.ParseInt(combined, 2, 32)
	return int(val)
}

func getMapMinAndMax(input map[Position]int) (min Position, max Position) {
	min = Position{0, 0}
	max = Position{0, 0}

	for k := range input {
		if k.x <= min.x {
			min.x = k.x
		}

		if k.y <= min.y {
			min.y = k.y
		}

		if k.x >= max.x {
			max.x = k.x
		}

		if k.y >= max.y {
			max.y = k.y
		}
	}

	return min, max
}

func extendBorder(imagePlane map[Position]int, min, max Position, amount, value int) {
	// top and bottom
	for i := min.y - amount; i < max.y+amount; i++ {

		imagePlane[Position{i, min.y - 1}] = value
		imagePlane[Position{i, min.y - 2}] = value
		imagePlane[Position{i, min.y - 3}] = value
		imagePlane[Position{i, max.y + 1}] = value
		imagePlane[Position{i, max.y + 2}] = value
		imagePlane[Position{i, max.y + 3}] = value
	}

	for i := min.x - amount; i < max.x+amount; i++ {
		imagePlane[Position{min.x - 1, i}] = value
		imagePlane[Position{min.x - 2, i}] = value
		imagePlane[Position{min.x - 3, i}] = value
		imagePlane[Position{max.x + 1, i}] = value
		imagePlane[Position{max.x + 2, i}] = value
		imagePlane[Position{max.x + 3, i}] = value
	}
}

func executeRuns(infiniteImagePlane map[Position]int, imageEnhancementAlgorithm []int, runs int, logging bool) int {
	// every single point in this is a light pixel. this will never include anything
	// regarding a dark pixel.
	for i := 0; i < runs; i++ {
		min, max := getMapMinAndMax(infiniteImagePlane)

		// if the index zero position is always on zero, then zero values will
		// never be converted to on, so we will never have infinite.
		infiniteModeEnabled := imageEnhancementAlgorithm[0] == 1

		// since all are off by default, infinite will only be on when we aer
		// on odd cases e.g. if i % 2 == 1
		isInfiniteTick := i%2 == 1

		// because it's all done INSTANTLY, we must have a update map
		// which is replacing the new map on completion.
		updatedMap := map[Position]int{}

		// if we are in the infinite mode, and  we are on the infinite
		//  tick lets go and increase our borders by 3x
		if infiniteModeEnabled {
			value := 0

			if isInfiniteTick {
				value = 1
			}

			// extend the border, so if we did have an infinite plane since
			// the infinite plane could apply changes to our main image if
			// the bits keep flipping.
			//
			// 3 since this is our overlapping amount that has no core
			// image involvement. Otherwise, it acts a big mad
			extendBorder(infiniteImagePlane, min, max, 3, value)
		}

		if logging {
			fmt.Println("------")
			min, max := getMapMinAndMax(infiniteImagePlane)
			printInfiniteGrid(infiniteImagePlane, min, max)
		}

		// iterate with the starting pixel being the bottom right and the
		// last pixel being checked being the top left
		for y := min.y - 1; y <= max.y+1; y++ {
			for x := min.x - 1; x <= max.x+1; x++ {
				updatedIndex := getPixelValueFromGrid(infiniteImagePlane, Position{x, y})

				if imageEnhancementAlgorithm[updatedIndex] == 1 {
					updatedMap[Position{x, y}] = imageEnhancementAlgorithm[updatedIndex]
				}
			}
		}

		infiniteImagePlane = updatedMap
	}

	if logging {
		fmt.Println("------")
		min, max := getMapMinAndMax(infiniteImagePlane)
		printInfiniteGrid(infiniteImagePlane, min, max)
		fmt.Println("------")
	}
	return len(infiniteImagePlane)
}

func main() {
	lines := helpers.ReadFileToText("./day20/input.txt")
	split := strings.Split(lines, "\n\n")

	// A two-dimensional grid of light pixels (#) and dark pixels (.)
	imageEnhancementAlgorithm := parseInputImage([]string{split[0]})[0]
	inputImage := parseInputImage(strings.Split(split[1], "\n"))

	infiniteImagePlane := map[Position]int{}

	for i := 0; i < len(inputImage); i++ {
		for j := 0; j < len(inputImage[i]); j++ {
			if inputImage[i][j] == 1 {
				infiniteImagePlane[Position{x: j, y: i}] = 1
			}
		}
	}

	fmt.Println("first", executeRuns(infiniteImagePlane, imageEnhancementAlgorithm, 2, false))
	fmt.Println("second", executeRuns(infiniteImagePlane, imageEnhancementAlgorithm, 50, false))
}

package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
	Next *Point
}

func diagonalShift(current *Point, end *Point) *Point {
	point := &Point{
		X:    current.X,
		Y:    current.Y,
		Next: nil,
	}

	if end.Y > current.Y {
		point.Y += 1
	} else {
		point.Y -= 1
	}

	if end.X > current.X {
		point.X += 1
	} else {
		point.X -= 1
	}

	return point
}

func horizontalShift(current *Point, end *Point) *Point {
	point := &Point{
		X:    current.X,
		Y:    current.Y,
		Next: nil,
	}

	if current.X == end.X {
		if end.Y > current.Y {
			point.Y += 1
		} else {
			point.Y -= 1
		}

		return point
	}

	if end.X > current.X {
		point.X += 1
	} else {
		point.X -= 1
	}

	return point
}

func getNext(current *Point, end *Point, horzMode bool) *Point {
	if horzMode {
		return horizontalShift(current, end)
	}

	return diagonalShift(current, end)
}

func output(points [][]*Point) int {
	pointCountMap := map[string]int{}

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[i]); j++ {
			point := points[i][j]

			name := fmt.Sprintf("%d-%d", point.X, point.Y)

			if _, ok := pointCountMap[name]; !ok {
				pointCountMap[name] = 0
			}

			pointCountMap[name] += 1
		}
	}
	count := 0

	for _, i := range pointCountMap {
		if i > 1 {
			count += 1
		}
	}

	return count
}

func run(path string, first bool) int {
	file := helpers.ReadFileToTextLines(path)

	pointDeltas := [][]*Point{}

	for _, position := range file {
		split := strings.Split(position, " -> ")

		// start
		leftSplit := strings.Split(split[0], ",")
		x1, _ := strconv.Atoi(leftSplit[0])
		y1, _ := strconv.Atoi(leftSplit[1])

		startPos := &Point{x1, y1, nil}

		// end
		rightSplit := strings.Split(split[1], ",")
		x2, _ := strconv.Atoi(rightSplit[0])
		y2, _ := strconv.Atoi(rightSplit[1])

		endPos := &Point{x2, y2, nil}
		next := startPos

		points := []*Point{startPos}

		horzMode := next.X == endPos.X || next.Y == endPos.Y

		if first && !horzMode {
			continue
		}

		for {
			next = getNext(points[len(points)-1], endPos, horzMode)
			points = append(points, next)

			if next.X == endPos.X && next.Y == endPos.Y {
				pointDeltas = append(pointDeltas, points)
				break
			}
		}
	}

	return output(pointDeltas)
}

func main() {
	path := "./day5/input.txt"

	fmt.Println("first", run(path, true))
	fmt.Println("second", run(path, false))
}

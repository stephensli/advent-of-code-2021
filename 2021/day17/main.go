package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers"
)

var (
	// Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if
	// it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
	xVelocityShift = 1
	// Due to gravity, the probe's y velocity decreases by 1.
	yVelocityShift = 1
)

type Target struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

//getMinxVelocity works by only incrementing by the velocity shift
// until we hit the Target.
//
// e.g 1+2+3+4+5=15
// 5 would hit 15
func getMinxVelocity(number int) (min int) {
	incline := 0

	for min < number {
		incline += xVelocityShift
		min += incline
	}

	return incline
}

func (t *Target) check() (madeTarget []*Rocket) {
	xmax := t.xMax
	xmin := getMinxVelocity(t.xMin)

	ymax := -t.yMin - 1
	ymin := t.yMin

	madeTarget = []*Rocket{}

	fmt.Println(xmin, xmax, ymin, ymax)

	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			r := &Rocket{
				target: t,
				x:      0,
				y:      0,
				velX:   x,
				velY:   y,
				deltaX: -xVelocityShift,
				deltaY: -yVelocityShift,
			}

			if r.madeTarget() {
				madeTarget = append(madeTarget, r)
			}
		}
	}

	return madeTarget
}

func partOne(madeTarget []*Rocket) int {
	maxY := 0

	for _, rocket := range madeTarget {
		if rocket.maxY > maxY {
			maxY = rocket.maxY
		}
	}

	return maxY
}

func parseTarget(input string) Target {
	var xMin, xMax, yMin, yMax int

	_, _ = fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &xMin, &xMax, &yMin, &yMax)
	return Target{xMin, xMax, yMin, yMax}
}

func main() {
	input := helpers.ReadFileToTextLines("./day17/input.txt")
	target := parseTarget(input[0])
	fmt.Printf("%v\n", target)

	madeTarget := target.check()

	fmt.Println(len(madeTarget), "made target")
	fmt.Printf("part one: %d\n", partOne(madeTarget))
	fmt.Printf("part two: %d\n", len(madeTarget))
}

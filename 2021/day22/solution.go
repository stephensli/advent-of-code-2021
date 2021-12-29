package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers/aoc"
	"github.com/stephensli/advent-of-code-2021/helpers/file"
	"github.com/stephensli/advent-of-code-2021/helpers/numbers"
	"github.com/stephensli/advent-of-code-2021/helpers/printers"
)

type Cuboid struct {
	X, Y, Z int
}

type Range struct {
	Min, Max, Abs int
}

type Step struct {
	onOff   bool
	X, Y, Z Range
}

// Calculate the area of a given cube.
func (b Step) volume() int {
	x := b.X.Max - b.X.Min + 1
	y := b.Y.Max - b.Y.Min + 1
	z := b.Z.Max - b.Z.Min + 1

	return x * y * z
}

func overLappingLine(min1, max1, min2, max2 int) (min int, max int) {
	maxMin := numbers.Max[int](min1, min2)
	minMax := numbers.Min[int](max1, max2)

	return maxMin, minMax
}

func getIntersectionOfCubes(a Step, steps []Step) int {
	result := 0

	for i, b := range steps {
		// get the min and max values between A min and B min, A max, B max. if
		// the max minus the minimum value is greater than zero, then the lines
		// must be overlapping. if all lines overlap in 1 way or another, then
		// we have an overlapping area we can compute.this overlap needs removing
		// since it cannot be turned on.
		minX, maxX := overLappingLine(a.X.Min, a.X.Max, b.X.Min, b.X.Max)
		minY, maxY := overLappingLine(a.Y.Min, a.Y.Max, b.Y.Min, b.Y.Max)
		minZ, maxZ := overLappingLine(a.Z.Min, a.Z.Max, b.Z.Min, b.Z.Max)

		if maxX-minX >= 0 && maxY-minY >= 0 && maxZ-minZ >= 0 {
			step := Step{
				X: Range{Min: minX, Max: maxX},
				Y: Range{Min: minY, Max: maxY},
				Z: Range{Min: minZ, Max: maxZ},
			}

			result += step.volume() - getIntersectionOfCubes(step, steps[i+1:])
		}
	}
	return result
}

func getCubesInRegion(region Step) (cubes []Cuboid) {
	for x := region.X.Min; x <= region.X.Max; x++ {
		for y := region.Y.Min; y <= region.Y.Max; y++ {
			for z := region.Z.Min; z <= region.Z.Max; z++ {
				cubes = append(cubes, Cuboid{
					x, y, z,
				})
			}
		}
	}

	return cubes
}

func partOne(steps []Step) {
	cubeState := map[Cuboid]bool{}

	for _, step := range steps {
		if step.X.Abs > 50 || step.Y.Abs > 50 || step.Z.Abs > 50 {
			continue
		}

		cubes := getCubesInRegion(step)

		// we can only turn on a cube if It's not already on, and if this cube
		// overlaps with an on cube then it's not turned on. while if we are
		// turning off cubes and over lap any cube that is on, it turns off
		for _, cube := range cubes {
			cubeState[cube] = step.onOff
		}

		// delete all cubes which are off for the next step
		for cuboid, b := range cubeState {
			if !b {
				delete(cubeState, cuboid)
			}
		}
	}

	aoc.PrintAnswer(1, len(cubeState))
}

func partTwo(steps []Step) {
	var processedSteps []Step
	result := 0

	for i := len(steps) - 1; i >= 0; i-- {
		step := steps[i]

		// working backwards means we can skip all the ones which are turned
		// off since the ones on are always going to be the most up-to-date
		// working backwards. If they were turned off and still off by the
		// start then nothing turned them back on.
		if step.onOff {
			// add the volume of the given cube plus any other lap
			// of any other cubes that have been processed.
			result += step.volume() - getIntersectionOfCubes(step, processedSteps)
		}

		// keep a list of what has been processed, so that the following steps
		// can reverse check to ensure all the intersections are removed.
		processedSteps = append(processedSteps, step)
	}

	aoc.PrintAnswer(2, result)
}

func parseInput(filePath string) []Step {
	var steps []Step

	for _, line := range file.ToTextLines(filePath) {
		step := Step{X: Range{}, Y: Range{}, Z: Range{}}

		var action string

		_, _ = fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &action, &step.X.Min, &step.X.Max,
			&step.Y.Min, &step.Y.Max, &step.Z.Min, &step.Z.Max)

		step.onOff = action == "on"

		// used in the first step to ensure we are not passing the 50 mark.
		step.X.Abs = numbers.Max[int](numbers.Abs[int](step.X.Min), numbers.Abs[int](step.X.Max))
		step.Y.Abs = numbers.Max[int](numbers.Abs[int](step.Y.Min), numbers.Abs[int](step.Y.Max))
		step.Z.Abs = numbers.Max[int](numbers.Abs[int](step.Z.Min), numbers.Abs[int](step.Z.Max))

		// if the step was the same action, go and merge them together, keeping the
		// latest values of all of them
		steps = append(steps, step)
	}

	printers.JsonPrint(steps, true)
	return steps
}

func main() {
	defer aoc.Setup(2021, 22)()

	steps := parseInput("./input.txt")

	partOne(steps)
	partTwo(steps)
}

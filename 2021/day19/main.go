package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers"
	"math"
	"strings"
)

func parseInput(inputLines string) []*Scanner {
	blocks := strings.Split(inputLines, "\n\n")
	var scanners []*Scanner

	for i, block := range blocks {
		var cords []ThreeDimensionPosition

		lines := strings.Split(block, "\n")[1:]

		for _, line := range lines {
			var x, y, z int64

			_, _ = fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			cords = append(cords, ThreeDimensionPosition{x, y, z})
		}

		newScanner := newScanner(i, ThreeDimensionPosition{0, 0, 0}, 0, cords)
		scanners = append(scanners, newScanner)
	}

	return scanners
}

func checkScannerIsNeighbour(scannerOne, scannerB *Scanner) *Scanner {
	// iterate over all possible rotations of the B beacons and pair them with
	// all possible A beacons. They are the same beacons, and we locate more
	// or equal to 12, then that's a neighbour.
	beaconsInA := scannerOne.GetBeaconsInWorld()

	// this could be speed up by only checking the size - 12 - 1, if no
	// matches in that amount, then there will be no point checking the
	// remaining 11 since it will never meet 12.
	for _, beaconInA := range beaconsInA {
		// for all possible rotations, we don't know the facing direction
		// for any of the scanners, so we must try to align with said
		// scanner and check all directions.
		for rotation := 0; rotation < 24; rotation += 1 {

			beaconsInB := scannerB.GetBeaconsInWorld()
			for _, beaconInB := range beaconsInB {
				// now get all determine a start position by diffing two
				// different beacons and use that to determine if we are
				// overlapping.
				newPosition := ThreeDimensionPosition{
					x: beaconInA.x - beaconInB.x,
					y: beaconInA.y - beaconInB.y,
					z: beaconInA.z - beaconInB.z,
				}

				count := 0

				// move scanner two to that position and gather all its beacons
				// again from the new position if 12 or more of these new beacons
				// are also matching scannerA beacons, then they are overlapping
				updatedScannerTwoPos := scannerB.setToLocation(newPosition)
				updatedScannerTwoBeacons := updatedScannerTwoPos.GetBeaconsInWorld()

				for _, updatedBeaconTwo := range updatedScannerTwoBeacons {
					// check that updatedBeaconTwo is within the beacons One
					// relative to world and increment, if count > 12 WIN.
					for _, beaconOneCheck := range beaconsInA {
						if beaconOneCheck.equal(updatedBeaconTwo) {
							count += 1
						}

						if count >= 12 {
							return updatedScannerTwoPos
						}
					}
				}
			}

			scannerB = scannerB.rotate()
		}

	}

	return nil
}

// try and find all the locations of the given scanners and set
// the location of the scanner. This is done by finding over
// lapping beacons
func locateScanners(scanners []*Scanner) (locatedScanners []*Scanner) {
	locatedScanners = []*Scanner{}
	locatedScannerMap := map[int]bool{}

	// for every single scanner that is located, push this onto
	// the queue to see if we can locate any of its other scanners
	// that are its neighbours.
	var checkScannerQueue []*Scanner

	locatedScanners = append(locatedScanners, scanners[0])
	checkScannerQueue = append(checkScannerQueue, scanners[0])
	locatedScannerMap[scanners[0].id] = true

	// remove first one
	scanners = scanners[1:]

	for len(checkScannerQueue) > 0 {
		scannerA := checkScannerQueue[0]
		checkScannerQueue = checkScannerQueue[1:]

		for _, scannerB := range scanners {
			// check for possible scanner neighbour
			locatedScannerB := checkScannerIsNeighbour(scannerA, scannerB)

			if locatedScannerB != nil {
				locatedScanners = append(locatedScanners, locatedScannerB)
				locatedScannerMap[locatedScannerB.id] = true

				checkScannerQueue = append(checkScannerQueue, locatedScannerB)
			}
		}

		// remove all scanners from located scanner map from the scanners array. Since
		// we know they are already in sync, we don't need to go do anything different.
		var newScanners []*Scanner

		for _, scanner := range scanners {
			if !locatedScannerMap[scanner.id] {
				newScanners = append(newScanners, scanner)
			}
		}

		scanners = newScanners
	}

	return locatedScanners
}

func partOne(locatedScanners []*Scanner) {
	uniqueBeaconsMap := map[ThreeDimensionPosition]bool{}
	values := []ThreeDimensionPosition{}

	for _, scanner := range locatedScanners {
		for _, position := range scanner.GetBeaconsInWorld() {
			if !uniqueBeaconsMap[position] {
				uniqueBeaconsMap[position] = true
				values = append(values, position)
			} else {
			}
		}
	}

	fmt.Println(len(uniqueBeaconsMap))
}

func partTwo(locatedScanners []*Scanner) {
	max := 0

	for _, scannerOne := range locatedScanners {
		for _, scannerTwo := range locatedScanners {
			if scannerOne.id != scannerTwo.id {
				maxScan := math.Abs(float64(scannerOne.position.x)-float64(scannerTwo.position.x)) +
					math.Abs(float64(scannerOne.position.y)-float64(scannerTwo.position.y)) +
					math.Abs(float64(scannerOne.position.z)-float64(scannerTwo.position.z))

				if int(maxScan) > max {
					max = int(maxScan)

				}
			}
		}

	}
	fmt.Println(max)
}

func main() {
	lines := helpers.ReadFileToText("./day19/input.txt")
	scanners := parseInput(lines)

	locatedScanners := locateScanners(scanners)
	partOne(locatedScanners)
	partTwo(locatedScanners)
}

package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/stephensli/aoc/helpers/aoc"
	"github.com/stephensli/aoc/helpers/file"
)

type direction string

var (
	north direction = "north"
	south direction = "south"
	east  direction = "east"
	west  direction = "west"
)

type Path struct {
	HitZero   bool
	HitGround bool

	nextX int
	nextY int

	fromDirection direction
	value         int

	boundary [][]int
}

type directions struct {
	from direction
	to   direction
	y, x int
}

type Tile struct {
	name       string
	directions map[direction]directions
}

type Position struct {
	tile  *Tile
	value int
}

func (p *Position) String() string {
	if p.tile == groundPipe || p.tile.name == "#" {
		return p.tile.name
	}

	return p.tile.name
	//return strconv.Itoa(p.value)
	//return p.tile.name
}

var (
	verticalPipe = &Tile{
		name: "|",
		directions: map[direction]directions{
			north: {from: north, to: south, y: 0, x: +1},
			south: {from: south, to: north, y: 0, x: -1},
		},
	}
	horizonalPipe = &Tile{
		name: "-",
		directions: map[direction]directions{
			west: directions{from: west, to: east, y: +1, x: 0},
			east: directions{from: east, to: west, y: -1, x: 0},
		},
	}
	lPipe = &Tile{
		name: "L",
		directions: map[direction]directions{
			north: directions{from: west, to: east, y: +1, x: 0},
			east:  directions{from: south, to: north, y: 0, x: -1},
		},
	}
	jPipe = &Tile{
		name: "J",
		directions: map[direction]directions{
			north: directions{from: east, to: west, y: -1, x: 0},
			west:  directions{from: south, to: north, y: 0, x: -1},
		},
	}
	sevenPipe = &Tile{
		name: "7",
		directions: map[direction]directions{
			west:  directions{from: north, to: south, y: 0, x: +1},
			south: directions{from: east, to: west, y: -1, x: 0},
		},
	}
	fPipe = &Tile{
		name: "F",
		directions: map[direction]directions{
			east:  directions{from: north, to: south, y: 0, x: +1},
			south: directions{from: west, to: east, y: +1, x: 0},
		},
	}
	groundPipe = &Tile{
		name: ".",
	}
	startPipe = &Tile{
		name: "S",
	}
)

var tileMap = map[string]*Tile{
	"|": verticalPipe,
	"-": horizonalPipe,
	"L": lPipe,
	"J": jPipe,
	"7": sevenPipe,
	"F": fPipe,
	".": groundPipe,
	"S": startPipe,
}

type followDir string

var (
	anyFollowDir   followDir = "any"
	leftFollowDir  followDir = "left"
	rightFollowDir followDir = "right"
)

var globalHistory = sync.Map{}

func checkDirectionsForBoundary(history map[string]bool, startX, startY int, dir direction, positions [][]*Position, followDirection followDir) (result bool, h map[string]bool) {

	if startX < 0 || startX > len(positions)-1 {
		return false, history
	}

	if startY < 0 || startY > len(positions[startX])-1 {
		return false, history
	}

	key := fmt.Sprintf("%d:%d", startX, startY)
	if value, ok := globalHistory.Load(key); ok {
		return value.(bool), history
	}

	if positions[startX][startY].value != -1 && positions[startX][startY].value != -2 {
		var left, right, current int = -99, -99, 0

		current = positions[startX][startY].value

		if dir == north || dir == south {
			if startY > 0 {
				left = positions[startX][startY-1].value
			}

			if startY < len(positions[startX])-1 {
				right = positions[startX][startY+1].value
			}

			// check up or down
		} else if dir == east || dir == west {
			if startX > 0 {
				left = positions[startX-1][startY].value
			}

			if startX < len(positions)-1 {
				right = positions[startX+1][startY].value
			}
		}

		if left == 0 && right == 0 {
			return false, history
		}

		leftValue := math.Abs(float64(current) - float64(left))
		rightValue := math.Abs(float64(current) - float64(right))

		if followDirection == anyFollowDir {
			if leftValue == 1 && rightValue == 1 {
				//fmt.Printf("pass - dir: %v, sX: %d, sY: %d, current: %v, left: %d, right: %d - history len : %d\n", dir, startX, startY, current, left, right, len(history))
				return true, history
			}
		}

		if followDirection == leftFollowDir {
			if leftValue == 1 {
				//	fmt.Printf("pass - dir: %v, sX: %d, sY: %d, current: %v, left: %d, right: %d - history len : %d\n", dir, startX, startY, current, left, right, len(history))
				return true, history
			}
		}

		if followDirection == rightFollowDir {
			if rightValue == 1 {
				// fmt.Printf("pass - dir: %v, sX: %d, sY: %d, current: %v, left: %d, right: %d - history len : %d\n", dir, startX, startY, current, left, right, len(history))
				return true, history
			}
		}

		if followDirection == anyFollowDir && leftValue != 1 && rightValue != 1 {
			followDirection = anyFollowDir
		} else if followDirection == anyFollowDir && leftValue == 1 {
			followDirection = rightFollowDir
		} else if followDirection == anyFollowDir && rightValue == 1 {
			followDirection = leftFollowDir
		}
		// fmt.Printf("dir: %v, sX: %d, sY: %d, current: %v, left: %d, right: %d, follow-dir: %s\n", dir, startX, startY, current, left, right, followDirection)
	}

	directionsValues := [][]int{}
	directions := []direction{}

	if followDirection == anyFollowDir {
		directions = append(directions, north, east, south, west)
		directionsValues = append(directionsValues, []int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1})
	} else {
		switch dir {
		case north:
			directions = append(directions, north)
			directionsValues = append(directionsValues, []int{-1, 0})
		case south:
			directions = append(directions, south)
			directionsValues = append(directionsValues, []int{1, 0})
		case east:
			directions = append(directions, east)
			directionsValues = append(directionsValues, []int{0, 1})
		case west:
			directions = append(directions, west)
			directionsValues = append(directionsValues, []int{0, -1})
		}
	}

	for i, v := range directionsValues {
		nextKey := fmt.Sprintf("%d:%d", startX+v[0], startY+v[1])
		if history[nextKey] {
			continue
		}

		history[key] = true

		if ok, _ := checkDirectionsForBoundary(
			history,
			startX+v[0],
			startY+v[1],
			directions[i],
			positions,
			followDirection,
		); !ok {
			return false, history
		}
	}

	return true, history
}

func getBoundingBoxCount(positions [][]*Position, path *Path) int {
	return 0
}

func startTraverseMap(positions [][]*Position, x, y int) (int, int) {
	positions[x][y].value = 0

	paths := []*Path{}

	if x-1 >= 0 && positions[x-1][y].tile != groundPipe {
		paths = append(paths, &Path{
			boundary:      [][]int{{x, y}},
			fromDirection: south,
			nextX:         x - 1,
			nextY:         y,
			value:         0,
		})
	}

	if x+1 < len(positions) && positions[x+1][y].tile != groundPipe {
		paths = append(paths, &Path{
			fromDirection: north,
			boundary:      [][]int{{x, y}},
			nextX:         x + 1,
			nextY:         y,
			value:         0,
		})
	}

	if y-1 >= 0 && positions[x][y-1].tile != groundPipe {
		paths = append(paths, &Path{
			boundary:      [][]int{{x, y}},
			fromDirection: east,
			nextX:         x,
			nextY:         y - 1,
			value:         0,
		})
	}

	if y+1 < len(positions[x]) && positions[x][y+1].tile != groundPipe {
		paths = append(paths, &Path{
			boundary:      [][]int{{x, y}},
			fromDirection: west,
			nextX:         x,
			nextY:         y + 1,
			value:         0,
		})
	}

	value := 1

	for {
		anyValid := false
		for _, v := range paths {
			if v.HitZero || v.HitGround {
				continue
			}

			next := positions[v.nextX][v.nextY]
			comingDirection := v.fromDirection

			setValue, hitZero, hitGround, validDirection := traverseMapAndSetValues(value, next.tile.directions[comingDirection].from, positions, v.nextX, v.nextY)

			v.HitZero = hitZero
			v.HitGround = hitGround

			valid := setValue || (!v.HitZero && !v.HitGround && validDirection)

			if setValue {
				v.value = value

			}

			if valid {
				v.boundary = append(v.boundary, []int{
					v.nextX,
					v.nextY,
				})

				v.nextX = v.nextX + next.tile.directions[comingDirection].x
				v.nextY = v.nextY + next.tile.directions[comingDirection].y
				v.fromDirection = next.tile.directions[comingDirection].from
				anyValid = true

			}
		}

		if !anyValid {
			break
		}

		value += 1
	}

	largest := 0
	var path *Path
	inLoopCount := 0
	for _, p := range paths {
		if p.HitZero && p.value > largest {
			largest = p.value
			path = p
		}
	}

	inLoopCount = getBoundingBoxCount(positions, path)
	return largest, inLoopCount
}

func traverseMapAndSetValues(value int, comingDirection direction, positions [][]*Position, x, y int) (setValue bool, hitZero bool, hitGround bool, validDirection bool) {
	validDirection = comingDirection != ""
	hitGround = positions[x][y].tile == groundPipe
	hitZero = positions[x][y].value == 0

	setValue = positions[x][y].value == -1 && !hitZero && !hitGround && validDirection

	if setValue {
		positions[x][y].value = value
	}

	return
}

func parser(lines [][]string) (startX, startY int, positions [][]*Position) {
	for i, row := range lines {
		positions = append(positions, make([]*Position, len(row)))

		for j, value := range row {
			if value == "S" {
				startX = i
				startY = j
			}

			positions[i][j] = &Position{tile: tileMap[value], value: -1}
		}
	}
	return
}

func main() {
	path, complete := aoc.Setup(2023, 10, false)
	defer complete()

	startX, startY, positions := parser(file.ToTextSplit(path, ""))

	partOne, partTwo := startTraverseMap(positions, startX, startY)

	aoc.PrintAnswer(1, partOne)
	aoc.PrintAnswer(2, partTwo)
}

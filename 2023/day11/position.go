package main

import (
	"strconv"

	"github.com/stephensli/aoc/helpers/algorithms"
)

type SpaceType int

var (
	EmptySpaceType  SpaceType = 0
	GalexySpaceType SpaceType = 1
	Filled          SpaceType = 9
)

type Position struct {
	Coords        algorithms.Coords
	PositionType  SpaceType
	PositionValue int
	SpaceWeight   int
}

func (p Position) Wall(direction algorithms.Direction) bool {
	return false
}

func (p Position) Position() algorithms.Coords {
	return p.Coords
}

func (p Position) Value() int {
	return p.PositionValue
}

func (p Position) Weight() int {
	return p.SpaceWeight
}

func (p Position) Empty() bool {
	return p.PositionType == EmptySpaceType
}

func (p Position) Galexy() bool {
	return p.PositionType == GalexySpaceType
}

func (p Position) IsPositionValue(value int) bool {
	return p.PositionValue == value
}

func (p Position) String() string {
	if p.PositionType == EmptySpaceType {
		return "."
	} else if p.PositionType == Filled {
		return "#"
	} else {
		return strconv.Itoa(p.PositionValue)
	}
}

package main

import (
	"github.com/life4/genesis/slices"
)

type Dimensions struct {
	Length, Width, Height int
}

type Present struct {
	Dimensions
}

func (p Present) GetSideAreas() []int {
	return []int{2 * p.Length * p.Width, 2 * p.Width * p.Height, 2 * p.Height * p.Length}
}

func (p Present) GetVolume() int {
	return p.Width * p.Height * p.Length
}

func (p Present) GetSurfaceArea() int {
	sides := p.GetSideAreas()

	return slices.Reduce(sides, 0, func(elm int, acc int) int {
		return elm + acc
	})
}

func (p Present) GetRibbonWithBow() int {
	sorted := slices.Sort([]int{p.Width, p.Height, p.Length})

	ribbonLength := (sorted[0] * 2) + (sorted[1] * 2)

	return ribbonLength + p.GetVolume()
}

func (p Present) GetWrappingPaperWithSlack() int {
	smallest, _ := slices.Min(p.GetSideAreas())
	return p.GetSurfaceArea() + (smallest / 2)
}

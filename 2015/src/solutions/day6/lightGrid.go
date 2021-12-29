package main

type Grid struct {
	on     int
	lights [][]int
}

func (g *Grid) ToggleLights(from, to Position, singleDigitBrightness bool) {
	for i := from.X; i <= to.X; i++ {
		for j := from.Y; j <= to.Y; j++ {
			if singleDigitBrightness {
				if g.lights[i][j] == 1 {
					g.lights[i][j] = 0
					g.on -= 1
				} else {
					g.lights[i][j] = 1
					g.on += 1
				}
			} else {
				g.lights[i][j] += 2
				g.on += 2
			}
		}
	}
}

func (g *Grid) TurnOnOffLights(onOff bool, from, to Position, singleDigitBrightness bool) {
	for i := from.X; i <= to.X; i++ {
		for j := from.Y; j <= to.Y; j++ {
			if singleDigitBrightness {
				if onOff && g.lights[i][j] == 0 {
					g.lights[i][j] = 1
					g.on += 1
				} else if !onOff && g.lights[i][j] == 1 {
					g.lights[i][j] = 0
					g.on -= 1
				}
			} else {
				if onOff {
					g.lights[i][j] += 1
					g.on += 1
				} else if !onOff && g.lights[i][j] > 0 {
					g.lights[i][j] -= 1
					g.on -= 1
				}
			}
		}
	}
}

func NewGrid(size int) *Grid {
	lights := make([][]int, size)

	for i := 0; i < size; i++ {
		lights[i] = make([]int, size)
	}

	return &Grid{
		on:     0,
		lights: lights,
	}
}

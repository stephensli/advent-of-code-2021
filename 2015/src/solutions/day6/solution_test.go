package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolutionPartOne(t *testing.T) {
	t.Run("toggle", func(t *testing.T) {
		t.Run("should turn all off or on", func(t *testing.T) {
			grid := NewGrid(10)

			grid.ToggleLights(Position{0, 0}, Position{9, 9}, true)
			assert.Equal(t, 10*10, grid.on)

			grid.ToggleLights(Position{0, 0}, Position{9, 9}, true)
			assert.Equal(t, 0, grid.on)
		})
	})

	t.Run("onOff", func(t *testing.T) {
		t.Run("on state should not change if all already on", func(t *testing.T) {
			grid := NewGrid(10)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, true)
			assert.Equal(t, 5*5, grid.on)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, true)
			assert.Equal(t, 5*5, grid.on)
		})

		t.Run("should turn off existing lights if already on", func(t *testing.T) {
			grid := NewGrid(10)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, true)
			assert.Equal(t, 5*5, grid.on)

			grid.TurnOnOffLights(false, Position{2, 1}, Position{4, 4}, true)
			assert.Equal(t, 13, grid.on)
		})

		t.Run("should not turn off existing lights if already off", func(t *testing.T) {
			grid := NewGrid(10)

			grid.TurnOnOffLights(false, Position{0, 0}, Position{4, 4}, true)
			assert.Equal(t, 0, grid.on)

		})
	})
}

func TestSolutionPartTwo(t *testing.T) {
	t.Run("toggle", func(t *testing.T) {
		t.Run("should increase brightness by 2x", func(t *testing.T) {
			grid := NewGrid(10)

			grid.ToggleLights(Position{0, 0}, Position{9, 9}, false)
			assert.Equal(t, 10*10*2, grid.on)

			grid.ToggleLights(Position{0, 0}, Position{9, 9}, false)
			assert.Equal(t, 10*10*4, grid.on)
		})
	})

	t.Run("onOff", func(t *testing.T) {
		t.Run("off should decrease by 1 until 0", func(t *testing.T) {
			grid := NewGrid(10)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, false)
			assert.Equal(t, 5*5, grid.on)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, false)
			assert.Equal(t, 5*5*2, grid.on)

			// attempt to decrease all to -1 (currently on two). should stay at zero.
			grid.TurnOnOffLights(false, Position{0, 0}, Position{4, 4}, false)
			grid.TurnOnOffLights(false, Position{0, 0}, Position{4, 4}, false)
			grid.TurnOnOffLights(false, Position{0, 0}, Position{4, 4}, false)

			assert.Equal(t, 0, grid.on)
		})

		t.Run("on should increase by 1", func(t *testing.T) {
			grid := NewGrid(10)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, false)
			assert.Equal(t, 5*5, grid.on)

			grid.TurnOnOffLights(true, Position{0, 0}, Position{4, 4}, false)
			assert.Equal(t, 5*5*2, grid.on)
		})

	})
}

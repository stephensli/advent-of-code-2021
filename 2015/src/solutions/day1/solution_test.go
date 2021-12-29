package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("should return correct target floor for input", func(t *testing.T) {
		var tests = []struct {
			input    []string
			expected int
		}{{
			input:    []string{"(", "(", ")", ")"},
			expected: 0,
		}, {
			input:    []string{"(", ")", ")", "(", ")"},
			expected: -1,
		}}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%v = %d", test.input, test.expected), func(t *testing.T) {
				floor, _ := determineFloor(test.input)
				assert.Equal(t, test.expected, floor)
			})
		}
	})

	t.Run("should return correct basement index for input", func(t *testing.T) {
		var tests = []struct {
			input    []string
			expected int
		}{{
			input: []string{"(", "(", ")", ")"},
			// zero meaning they did not enter the basement.
			expected: math.MinInt,
		}, {
			input:    []string{"(", ")", ")", "(", ")"},
			expected: 3,
		}, {
			input:    []string{"(", ")", "(", ")", ")"},
			expected: 5,
		}}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%v = %d", test.input, test.expected), func(t *testing.T) {
				_, baseIndex := determineFloor(test.input)
				assert.Equal(t, test.expected, baseIndex)
			})
		}
	})
}

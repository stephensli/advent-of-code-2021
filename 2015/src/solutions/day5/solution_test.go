package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNicePartTwo(t *testing.T) {
	t.Run("contains a pair of any two letters that appear twice (ignore overlapping)", func(t *testing.T) {
		var tests = []struct {
			input    []string
			expected bool
		}{{
			// this is not nice since it only contains 1 valid pair of
			// two letters that appear next to each other.
			input:    []string{"a", "a", "a"},
			expected: false,
		}, {
			// this is valid, but it should only return that
			// two have been found and not three.
			input:    []string{"a", "a", "a", "a"},
			expected: true,
		}, {
			input:    []string{"q", "j", "h", "v", "h", "t", "z", "x", "z", "q", "q", "j", "k", "m", "p", "b"},
			expected: true,
		}}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%v = %v", test.input, test.expected), func(t *testing.T) {
				answer := isNicePartTwo(test.input)
				assert.Equal(t, test.expected, answer)
			})
		}
	})
}

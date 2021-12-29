package main

import (
	"encoding/json"
	"github.com/stephensli/advent-of-code-2021/helpers/file"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSeaCucumberExample(t *testing.T) {
	t.Run("example should complete after 58 steps", func(t *testing.T) {
		lines := file.ToTextSplit("./input.example.txt", "")
		input := parseInput(lines)

		expectedOutContent := [][]SeaCucumber{}
		fileContent, _ := ioutil.ReadFile("./out.example.json")
		_ = json.Unmarshal(fileContent, &expectedOutContent)

		resp, changes := StepCucumbers(input)

		count := 1
		for {
			if changes == 0 || count > 58 {
				break
			}

			count += 1

			resp, changes = StepCucumbers(resp)
		}

		assert.Equal(t, count, 58)
		assert.Equal(t, resp, expectedOutContent)
	})
}

func TestSeaCucumberCount(t *testing.T) {
	t.Run("should return the count of the changes", func(t *testing.T) {
		input := []string{"vv.", "..v", ".v."}
		expected := [][]SeaCucumber{{0, 0, 0}, {2, 2, 0}, {0, 2, 2}}

		seaCucumbers := parseInput([][]string{
			strings.Split(input[0], ""),
			strings.Split(input[1], ""),
			strings.Split(input[2], ""),
		})

		resp, count := StepCucumbers(seaCucumbers)

		assert.Equal(t, expected, resp)
		assert.Equal(t, count, 3)
	})

	t.Run("Should return zero on no changes & boards should equal", func(t *testing.T) {
		input := []string{".v.", ".v.", ".v."}
		expected := [][]SeaCucumber{{0, 2, 0}, {0, 2, 0}, {0, 2, 0}}

		seaCucumbers := parseInput([][]string{
			strings.Split(input[0], ""),
			strings.Split(input[1], ""),
			strings.Split(input[2], ""),
		})

		resp, count := StepCucumbers(seaCucumbers)

		assert.Equal(t, expected, resp)
		assert.Equal(t, count, 0)
	})
}

func TestEatFacing(t *testing.T) {
	t.Run("If east facing are grouped, only the right most will move", func(t *testing.T) {
		var tests = []struct {
			input    string
			expected [][]SeaCucumber
		}{
			{
				input:    "...>>>>>...",
				expected: [][]SeaCucumber{{0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0}},
			},
			{
				input:    "...>>>>.>..",
				expected: [][]SeaCucumber{{0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0}},
			},
			{
				input:    "...>>>.>.>.",
				expected: [][]SeaCucumber{{0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1}},
			},
		}

		for _, test := range tests {
			input := parseInput([][]string{strings.Split(test.input, "")})
			resp, _ := StepCucumbers(input)

			assert.Equal(t, test.expected, resp)
		}
	})

	t.Run("if on right boarder it should appear on the left side on next step", func(t *testing.T) {
		input := ".>.>"
		expected := [][]SeaCucumber{{1, 0, 1, 0}}

		seaCucumbers := parseInput([][]string{strings.Split(input, "")})
		resp, _ := StepCucumbers(seaCucumbers)

		assert.Equal(t, expected, resp)
	})

	t.Run("SC cannot wrap around if left most is filled.", func(t *testing.T) {
		input := ">>.>"
		expected := [][]SeaCucumber{{1, 0, 1, 1}}

		seaCucumbers := parseInput([][]string{strings.Split(input, "")})
		resp, _ := StepCucumbers(seaCucumbers)

		assert.Equal(t, expected, resp)
	})
}

func TestSouthFacing(t *testing.T) {
	t.Run("If south facing are grouped, only the right most will move", func(t *testing.T) {
		var tests = []struct {
			input    []string
			expected [][]SeaCucumber
		}{
			{
				input:    []string{".v.", ".v."},
				expected: [][]SeaCucumber{{0, 2, 0}, {0, 2, 0}},
			},
			{
				input:    []string{"vv.", ".v."},
				expected: [][]SeaCucumber{{0, 2, 0}, {2, 2, 0}},
			},
		}

		for _, test := range tests {
			input := parseInput([][]string{
				strings.Split(test.input[0], ""),
				strings.Split(test.input[1], ""),
			})

			resp, _ := StepCucumbers(input)

			assert.Equal(t, test.expected, resp)
		}
	})

	t.Run("if on bottom boarder it should appear on the top side on next step", func(t *testing.T) {
		input := []string{"v.v", ".v."}
		expected := [][]SeaCucumber{{0, 2, 0}, {2, 0, 2}}

		seaCucumbers := parseInput([][]string{strings.Split(input[0], ""), strings.Split(input[1], "")})
		resp, _ := StepCucumbers(seaCucumbers)

		assert.Equal(t, expected, resp)
	})

	t.Run("SC cannot wrap around if top most is filled.", func(t *testing.T) {
		input := []string{".v.", "...", ".v."}
		expected := [][]SeaCucumber{{0, 0, 0}, {0, 2, 0}, {0, 2, 0}}

		seaCucumbers := parseInput([][]string{
			strings.Split(input[0], ""),
			strings.Split(input[1], ""),
			strings.Split(input[2], ""),
		})

		resp, _ := StepCucumbers(seaCucumbers)

		assert.Equal(t, expected, resp)
	})
}

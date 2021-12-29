package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{3, 3, 6},
		{2, -3, -1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", test.a, test.b, test.expected), func(t *testing.T) {
			ok := add(&test.a, &test.b)

			assert.Equal(t, test.a, test.expected)
			assert.Equal(t, ok, true)
		})
	}
}

func TestEql(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 0},
		{3, 3, 1},
		{2, -3, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d == %d = %d", test.a, test.b, test.expected), func(t *testing.T) {
			ok := eql(&test.a, &test.b)

			assert.Equal(t, test.a, test.expected)
			assert.Equal(t, ok, true)
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{3, 3, 9},
		{2, -3, -6},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d * %d = %d", test.a, test.b, test.expected), func(t *testing.T) {
			ok := mul(&test.a, &test.b)

			assert.Equal(t, test.a, test.expected)
			assert.Equal(t, ok, true)
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{6, 3, 2},
		{3, 3, 1},
		{10, 2, 5},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d / %d = %d", test.a, test.b, test.expected), func(t *testing.T) {
			ok := div(&test.a, &test.b)

			assert.Equal(t, test.a, test.expected)
			assert.Equal(t, ok, true)
		})
	}

	t.Run("fail if B=0", func(t *testing.T) {
		a := 5
		b := 0

		ok := div(&a, &b)
		assert.Equal(t, ok, false)
		assert.Equal(t, a, 5)
		assert.Equal(t, b, 0)
	})
}

func TestMod(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 2},
		{3, 3, 0},
		{2, 13, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d mod %d = %d", test.a, test.b, test.expected), func(t *testing.T) {
			ok := mod(&test.a, &test.b)

			assert.Equal(t, test.a, test.expected)
			assert.Equal(t, ok, true)
		})
	}

	t.Run("fail if a<0", func(t *testing.T) {
		a := -1
		b := 5

		ok := mod(&a, &b)

		assert.Equal(t, ok, false)
		assert.Equal(t, a, -1)
		assert.Equal(t, b, 5)
	})

	t.Run("fail if b<0", func(t *testing.T) {
		a := 10
		b := -1

		ok := mod(&a, &b)

		assert.Equal(t, ok, false)
		assert.Equal(t, a, 10)
		assert.Equal(t, b, -1)
	})

	t.Run("fail if b=0", func(t *testing.T) {
		a := 10
		b := 0

		ok := mod(&a, &b)

		assert.Equal(t, ok, false)
		assert.Equal(t, a, 10)
		assert.Equal(t, b, 0)
	})

}

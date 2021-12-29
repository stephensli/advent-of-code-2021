package main

import (
	"fmt"
	"github.com/stephensli/advent-of-code-2021/helpers"
	"math"
	"strconv"
)

type SnailNode struct {
	value       *int64
	left, right *SnailNode
}

func (s *SnailNode) hasValue() bool {
	return s.value != nil
}

func (s *SnailNode) magnitude() int64 {
	if s.value != nil {
		return *s.value
	}

	return 3*s.left.magnitude() + 2*s.right.magnitude()
}

func (s *SnailNode) makeSplits() bool {
	if s.hasValue() && *s.value >= 10 {
		left := int64(math.Floor(float64(*s.value) / 2))
		right := int64(math.Ceil(float64(*s.value) / 2))

		s.left = &SnailNode{value: &left}
		s.right = &SnailNode{value: &right}
		s.value = nil
		return true
	}

	if s.hasValue() {
		return false
	}

	if s.left.makeSplits() {
		return true
	}

	return s.right.makeSplits()
}

func (s *SnailNode) makeExplosion(height int) (left, right int64) {
	if !s.hasValue() && s.left.hasValue() && s.right.hasValue() && height > 3 {
		left, right := s.left.value, s.right.value

		emptyZero := int64(0)
		s.value = &emptyZero
		s.left = nil
		s.right = nil
		return *left, *right
	}

	if s.hasValue() {
		return 0, 0
	}

	ll, rl := s.left.makeExplosion(height + 1)

	if rl > 0 {
		child := s.right

		for !child.hasValue() {
			child = child.left
		}

		*child.value += rl
	}

	lr, rr := s.right.makeExplosion(height + 1)

	if lr > 0 {
		child := s.left

		for !child.hasValue() {
			child = child.right
		}

		*child.value += lr
	}

	return ll, rr
}

func parse(input string) *SnailNode {

	root := &SnailNode{}
	stack := []*SnailNode{}

	stack = append(stack, root)

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '[':
			left := &SnailNode{}
			stack[len(stack)-1].left = left
			stack = append(stack, left)
			break
		case ']':
			stack = stack[:len(stack)-1]
			break
		case ',':
			stack = stack[:len(stack)-1]
			right := &SnailNode{}

			stack[len(stack)-1].right = right
			stack = append(stack, right)
			break
		default:
			val, _ := strconv.ParseInt(string(input[i]), 10, 32)
			stack[len(stack)-1].value = &val
		}
	}

	return root
}

func reduce(node *SnailNode) {
	for {
		node.makeExplosion(0)
		change := node.makeSplits()

		if !change {
			break
		}
	}
}

func addSnailNodes(a, b *SnailNode) *SnailNode {
	newNode := &SnailNode{}

	newNode.left = a
	newNode.right = b

	// reduce
	reduce(newNode)

	return newNode
}

func main() {
	input := helpers.ReadFileToTextLines("./day18/input.txt")
	inputSecond := helpers.ReadFileToTextLines("./day18/input.txt")

	first := input[0]
	input = input[1:]

	root := parse(first)

	// for each value, the two nodes need adding after they add they
	// must be reduced which means explode, split, If no split, then
	// complete otherwise repeat.
	for {
		if len(input) == 0 {
			break
		}

		next := input[0]
		input = input[1:]

		root = addSnailNodes(root, parse(next))
	}

	max := int64(0)

	for i := 0; i < len(inputSecond); i++ {
		for j := 0; j < len(inputSecond); j++ {
			first := parse(inputSecond[i])
			second := parse(inputSecond[j])

			mag := addSnailNodes(first, second).magnitude()
			if mag > max {
				max = mag
			}

		}

	}

	fmt.Println("first", root.magnitude())
	fmt.Println("second", max)

}

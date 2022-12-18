package queue

import (
	"golang.org/x/exp/constraints"
)

type Stack[T constraints.Ordered] struct {
	stack []T
}

func (pq *Stack[T]) Peak() T { return pq.stack[len(pq.stack)-1] }

func (pq *Stack[T]) Value(index int) T { return pq.stack[index] }

func (pq *Stack[T]) Len() int { return len(pq.stack) }

func (pq *Stack[T]) Less(i, j int) bool {
	return pq.stack[i] > pq.stack[j]
}

func (pq *Stack[T]) Swap(i, j int) {
	pq.stack[i], pq.stack[j] = pq.stack[j], pq.stack[i]
}

func (pq *Stack[T]) Push(x ...T) {
	pq.stack = append(pq.stack, x...)
}

func (pq *Stack[T]) Pop() T {
	var noop T
	old := pq.stack
	n := len(old)
	item := old[n-1]
	old[n-1] = noop
	pq.stack = old[0 : n-1]
	return item
}

package queue

type Queue[T any] struct {
	stack []T
}

func (pq *Queue[T]) Peak() T { return pq.stack[len(pq.stack)-1] }

func (pq *Queue[T]) Value(index int) T { return pq.stack[index] }

func (pq *Queue[T]) Len() int { return len(pq.stack) }

func (pq *Queue[T]) Swap(i, j int) {
	pq.stack[i], pq.stack[j] = pq.stack[j], pq.stack[i]
}

func (pq *Queue[T]) Push(x ...T) {
	pq.stack = append(pq.stack, x...)
}

func (pq *Queue[T]) Pop() T {
	x := pq.stack[0]
	pq.stack = pq.stack[1:]
	return x
}

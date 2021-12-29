package queue

// An Item is something we manage in a priority queue.
type Item struct {
	Value    any // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

package pkg

import (
	"context"
	"monago/utils"
)

// Queue is a MPMS (multiple producer, multiple subscriber) queue that is
// thread-safe and does not contain unsafe code. The only catch is that the
// queue has a fixed size. Once the queue is full, insertions will block
// forever until the queue has room.
type Queue[T any] struct {
	data chan T
}

// NewQueue creates a Queue of a given size
func NewQueue[T any](size int) Queue[T] {
	return Queue[T]{
		data: make(chan T, size),
	}
}

// Push adds a value to the queue, maybe blocking forever if the queue is full.
func (q Queue[T]) Push(val T) {
	q.data <- val
}

// Pop removes a value from the queue, maybe blocking forever if the queue is empty.
func (q Queue[T]) Pop() T {
	return <-q.data
}

// TryPop tries to pop from a queue. If it stalls long enough for the context
// to time out, it panics.
func (q Queue[T]) TryPop(ctx context.Context) (T, bool) {
	select {
	case val := <-q.data:
		return val, true
	case <-ctx.Done():
		return utils.Zero[T](), false
	}
}

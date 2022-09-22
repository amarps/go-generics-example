package queue

import "errors"

type Queue[T any] struct {
	collections []T
}

var (
	ErrEmptyDequeue = errors.New("cannot dequeue an empty queue")
)

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (Q *Queue[T]) Clear() {
	Q.collections = []T{}
}

func (Q *Queue[T]) IsEmpty() bool {
	return Q.Len() == 0
}

func (Q *Queue[T]) Len() int {
	return len(Q.collections)
}

func (Q *Queue[T]) Enqueue(v T) {
	Q.collections = append(Q.collections, v)
}

func (Q *Queue[T]) Dequeue() T {
	if Q.Len() == 0 {
		panic(ErrEmptyDequeue)
	}
	res := Q.collections[0]
	Q.collections = Q.collections[1:]
	return res
}

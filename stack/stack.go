package stack

import "errors"

type Stack[T any] struct {
	collections []T
}

var (
	ErrEmptyPop = errors.New("cannot pop an empty stack")
)

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (S *Stack[T]) Clear() {
	S.collections = []T{}
}

func (S *Stack[T]) IsEmpty() bool {
	return S.Len() == 0
}

func (S *Stack[T]) Len() int {
	return len(S.collections)
}

func (S *Stack[T]) Push(v T) {
	S.collections = append(S.collections, v)
}

func (S *Stack[T]) Pop() T {
	len := S.Len()
	if len == 0 {
		panic(ErrEmptyPop)
	}
	res := S.collections[len-1]
	S.collections = S.collections[:len-1]
	return res
}

package stack_test

import (
	"testing"

	"github.com/amarps/go-generics-example/stack"
)

func TestStack(t *testing.T) {
	t.Run("simple push and pop success", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(5)
		s.Push(6)
		s.Push(7)
		if s.Len() != 3 {
			t.Errorf("Expected %d, got: %d", 1, s.Len())
		}

		item := s.Pop() == 7
		if !item {
			t.Errorf("Expected %t, got: %t", true, item)
		}

		item = s.Pop() == 6
		if !item {
			t.Errorf("Expected %t, got: %t", true, item)
		}

		item = s.Pop() == 5
		if !item {
			t.Errorf("Expected %t, got: %t", true, item)
		}

		if !s.IsEmpty() {
			t.Errorf("Expected %t, got: %t", true, s.IsEmpty())
		}
	})

	t.Run("pop on empty stack", func(t *testing.T) {
		s := stack.New[int]()
		assertPanics(t, stack.ErrEmptyPop.Error(), func() {
			s.Pop()
		})
	})

	t.Run("stack length should be 0 when created", func(t *testing.T) {
		s := stack.New[int]()
		if s.Len() != 0 {
			t.Error("stack len is not zero when created")
		}
	})

	t.Run("is empty should be true when created", func(t *testing.T) {
		s := stack.New[int]()
		if !s.IsEmpty() {
			t.Error("stack len is not zero when created")
		}
	})

	t.Run("test stack has correct length", func(t *testing.T) {
		s := stack.New[int]()
		if s.Len() != 0 {
			t.Error("stack len is not zero when created")
		}

		for length := 0; length < 1000; length++ {
			s.Push(length)
			if s.Len() != length+1 {
				t.Error("Pushing", length, "has length", s.Len())
			}
		}

		for length := 0; length < 1000; length++ {
			s.Pop()
			if s.Len() != 1000-length-1 {
				t.Error("Poping", length, "has length", s.Len())
			}
		}
	})

	t.Run("test stack item is correct", func(t *testing.T) {
		s := stack.New[int]()
		if s.Len() != 0 {
			t.Error("stack len is not zero when created")
		}

		elements := make(map[int]int)

		for length := 0; length < 1000; length++ {
			item := length + 2
			elements[length] = item
			s.Push(item)
		}

		for length := 999; !s.IsEmpty(); length-- {
			item := s.Pop()
			if item != elements[length] {
				t.Error("Poping", elements[length], "has item", item)
			}
		}
	})
}

func BenchmarkStack(b *testing.B) {
	s := stack.New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; s.IsEmpty(); i++ {
		s.Pop()
	}
}

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: expected panic", name)
		}
	}()

	f()
}

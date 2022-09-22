package queue_test

import (
	"testing"

	"github.com/amarps/go-generics-example/queue"
)

func TestQueue(t *testing.T) {
	t.Run("simple enqueue and dequeue success", func(t *testing.T) {
		q := queue.New[int]()
		q.Enqueue(5)
		q.Enqueue(6)
		q.Enqueue(7)
		if q.Len() != 3 {
			t.Errorf("Expected %d, got: %d", 1, q.Len())
		}

		item := q.Dequeue() == 5
		if !item {
			t.Errorf("Expected %t, got: %t", true, item)
		}

		item = q.Dequeue() == 6
		if !item {
			t.Errorf("Expected %t, got: %t", true, item)
		}

		item = q.Dequeue() == 7
		if !item {
			t.Errorf("Expected %t, got: %t", true, item)
		}

		if !q.IsEmpty() {
			t.Errorf("Expected %t, got: %t", true, q.IsEmpty())
		}
	})

	t.Run("dequeue on empty queue", func(t *testing.T) {
		q := queue.New[int]()
		assertPanics(t, queue.ErrEmptyDequeue.Error(), func() {
			q.Dequeue()
		})
	})

	t.Run("queue length should be 0 when created", func(t *testing.T) {
		q := queue.New[int]()
		if q.Len() != 0 {
			t.Error("queue len is not zero when created")
		}
	})

	t.Run("is empty should be true when created", func(t *testing.T) {
		q := queue.New[int]()
		if !q.IsEmpty() {
			t.Error("queue len is not zero when created")
		}
	})

	t.Run("test queue has correct length", func(t *testing.T) {
		q := queue.New[int]()
		if q.Len() != 0 {
			t.Error("queue len is not zero when created")
		}

		for length := 0; length < 1000; length++ {
			q.Enqueue(length)
			if q.Len() != length+1 {
				t.Error("enqueueing", length, "has length", q.Len())
			}
		}

		for length := 0; length < 1000; length++ {
			q.Dequeue()
			if q.Len() != 1000-length-1 {
				t.Error("dequeueing", length, "has length", q.Len())
			}
		}
	})

	t.Run("test queue item is correct", func(t *testing.T) {
		q := queue.New[int]()
		if q.Len() != 0 {
			t.Error("queue len is not zero when created")
		}

		elements := make(map[int]int)

		for length := 0; length < 1000; length++ {
			item := length + 2
			elements[length] = item
			q.Enqueue(item)
		}

		for length := 0; !q.IsEmpty(); length++ {
			item := q.Dequeue()
			if item != elements[length] {
				t.Error("dequeueing", elements[length], "has item", item)
			}
		}
	})
}

func BenchmarkQueue(b *testing.B) {
	q := queue.New[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; q.IsEmpty(); i++ {
		q.Dequeue()
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

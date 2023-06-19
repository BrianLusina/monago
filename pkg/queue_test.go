package pkg

import (
	"context"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int](5)
	for i := range make([]struct{}, 5) {
		q.Push(i)
	}

	for range make([]struct{}, 5) {
		t.Log(q.Pop())
	}

	defer func() {
		if r := recover(); r != nil {
			if sr, ok := r.(string); ok {
				t.Log(sr)
				return
			}
			panic(r)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	q.Push(1)
	t.Log(q.TryPop(ctx))
	q.TryPop(ctx)
}

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue[int](5)
		for i := range make([]struct{}, 5) {
			q.Push(i)
		}

		for range make([]struct{}, 5) {
			b.Log(q.Pop())
		}

		defer func() {
			if r := recover(); r != nil {
				if sr, ok := r.(string); ok {
					b.Log(sr)
					return
				}
				panic(r)
			}
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		q.Push(1)
		b.Log(q.TryPop(ctx))
		q.TryPop(ctx)
	}
}

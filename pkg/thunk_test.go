package pkg

import "testing"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func TestRecurFib(t *testing.T) {
	t.Log(fib(40))
}

func BenchmarkRecurFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(fib(40))
	}
}

func TestThunkFib(t *testing.T) {
	cache := make([]*Thunk[int], 41)

	fib := func(n int) int {
		return cache[n-1].Force() + cache[n-2].Force()
	}

	for i := range cache {
		i := i
		cache[i] = NewThunk(func() int { return fib(i) })
	}

	cache[0].option.Set(0)
	cache[1].option.Set(1)

	t.Log(cache[40].Force())
}

func TestMemoizedFib(t *testing.T) {
	mem := map[int]int{
		0: 0,
		1: 1,
	}

	var fib func(int) int
	fib = func(n int) int {
		if result, ok := mem[n]; ok {
			return result
		}

		result := fib(n-1) + fib(n-2)
		mem[n] = result
		return result
	}

	t.Log(fib(40))
}

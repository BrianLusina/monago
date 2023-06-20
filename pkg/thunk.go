package pkg

// Thunk contains a partially computed value. This thunk is also lazy because
// it will only run the underlying thunked action _once_ and then return the
// result of that action, caching it for further use.
//
// Consider the following JavaScript:
//
//	let add = (x, y) => x + y;
//	let addCurr = (x) => (y) => x + y;
//	console.log(add(2, 2)); // 4
//	console.log(addCurr(2)(2)); // 4
//	let addTwo = addCurr(2); // (y) => 2 + y;
//
// In this example, `addTwo` is a thunk that contains a partially applied addCurr
// invocation.
type Thunk[T any] struct {
	// doer is an action to be thunked
	doer func() T

	// option is the cache for complete data
	option *Option[T]
}

// NewThunk creates a new thunk
func NewThunk[T any](action func() T) *Thunk[T] {
	return &Thunk[T]{
		doer:   action,
		option: NewOption[T](),
	}
}

// Force evaluates a Thunk's action if it needs to, otherwise it returns the previously evaluated value
func (t Thunk[T]) Force() T {
	if t.option.IsSome() {
		return t.option.Yank()
	}
	t.option.Set(t.doer())
	return t.option.Yank()
}

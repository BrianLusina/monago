package pkg

import (
	"errors"
	"monago/utils"
)

var (
	ErrOptionIsNone = errors.New("monago: Option[T] has no value")
)

// Option is a container that might contain a value
type Option[T any] struct {
	val *T
}

// NewOption creates a new Option with the given type. All options start out as nil
func NewOption[T any]() *Option[T] {
	return &Option[T]{}
}

// Take a value out of the option if it exists, returns an Error if option is nil
func (o Option[T]) Take() (T, error) {
	if o.IsNone() {
		return utils.Zero[T](), ErrOptionIsNone
	}

	return *o.val, nil
}

// Set sets a value fo the option
func (o *Option[T]) Set(val T) {
	o.val = &val
}

// Clear removes the value from an option
func (o *Option[T]) Clear() {
	o.val = nil
}

// IsSome returns true if the option has some value in it
func (o Option[T]) IsSome() bool {
	return o.val != nil
}

// IsNone returns true if the option has no value in it
func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

// Yank will pull a value out of an option panicking if it does not exist
func (o Option[T]) Yank() T {
	if o.IsNone() {
		panic("monago: Yank on a None Option")
	}

	return *o.val
}

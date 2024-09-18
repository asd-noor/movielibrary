package types

type Result[T any] struct {
	value T
	err   error
}

// New returns a pointer to Result after constructing it
func NewResult[T any](value T) *Result[T] {
	return &Result[T]{
		value: value,
		err:   nil,
	}
}

func (r *Result[T]) Ok() bool {
	return r.err == nil
}

// SetErr sets the error field and returns the pointer
func (r *Result[T]) SetErr(err error) *Result[T] {
	r.err = err
	return r
}

// Unwrap destructs the Result, returning underlying value and error
func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

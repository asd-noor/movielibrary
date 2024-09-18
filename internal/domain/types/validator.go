package types

type Validator[T any] struct {
	Fn     func() bool
	ErrMsg string
}

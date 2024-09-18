package types

import (
	"encoding/json"
	"fmt"
)

type Option[T any] struct {
	value  T
	exists bool
}

func Some[T any](value T) *Option[T] {
	return &Option[T]{value: value, exists: true}
}

func None[T any]() *Option[T] {
	return &Option[T]{exists: false}
}

func (o *Option[T]) Get() T {
	if !o.exists {
		panic("cannot get value from none")
	}

	return o.value
}

func (o *Option[T]) GetOrDefault(defaultValue T) T {
	if !o.exists {
		return defaultValue
	}

	return o.value
}

func (o *Option[T]) IsNone() bool {
	return !o.exists
}

func (o *Option[_]) String() string {
	if o.exists {
		return fmt.Sprintf("Some(%v)", o.value)
	}

	return "None"
}

func (o *Option[_]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.exists = false

		return nil
	}

	err := json.Unmarshal(data, &o.value)
	if err != nil {
		return err
	}

	o.exists = true

	return nil
}

func (o Option[_]) MarshalJSON() ([]byte, error) {
	if o.exists {
		return json.Marshal(o.value)
	} else {
		return []byte("null"), nil
	}
}

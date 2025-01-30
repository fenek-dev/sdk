package util

// COPIED FROM github.com/samber/lo
// "A little coping better than a little dependencing"

// Empty returns an empty value.
func Empty[T any]() T {
	var zero T
	return zero
}

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// FromPtr returns the pointer value or empty.
func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}

	return *x
}

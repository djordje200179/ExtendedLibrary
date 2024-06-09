package misc

// Cloner is a type that can clone itself.
type Cloner[T any] interface {
	// Clone returns a copy of the object.
	Clone() T
}

package comparison

// Comparator is a function that compares two values of the same type.
type Comparator[T any] func(first, second T) int

const (
	FirstSmaller  int = -1
	Equal             = 0
	SecondSmaller     = 1

	FirstBigger  = SecondSmaller
	SecondBigger = FirstSmaller
)

// Reverse returns a new Comparator that
// compares two values in the opposite order.
func (comp Comparator[T]) Reverse() Comparator[T] {
	return func(first, second T) int { return -comp(first, second) }
}

// Less returns a function that Comparator two values
// and returns true if the first value is smaller than the second.
func (comp Comparator[T]) Less() func(first, second T) bool {
	return func(first, second T) bool { return comp(first, second) == FirstSmaller }
}

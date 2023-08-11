package comparison

type Comparator[T any] func(first, second T) int

const (
	FirstSmaller  int = -1
	Equal             = 0
	SecondSmaller     = 1

	FirstBigger  = SecondSmaller
	SecondBigger = FirstSmaller
)

func (comparator Comparator[T]) Reverse() Comparator[T] {
	return func(first, second T) int {
		return -comparator(first, second)
	}
}

func (comparator Comparator[T]) Less() func(first, second T) bool {
	return func(first, second T) bool {
		return comparator(first, second) == -1
	}
}

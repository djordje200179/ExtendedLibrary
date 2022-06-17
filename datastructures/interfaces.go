package datastructures

type Iterator[T any] interface {
	Valid() bool
	Move()

	Get() T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Collection[K comparable, V any] interface {
	Size() int

	Get(K) V
	Set(K, V)
}

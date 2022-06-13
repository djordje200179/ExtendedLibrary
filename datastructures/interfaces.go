package datastructures

type Iterator[T any] interface {
	IsValid() bool
	Move()

	Get() T
}

type Collection[K comparable, V any] interface {
	Size() int

	Get(K) V
	Set(K, V)
}

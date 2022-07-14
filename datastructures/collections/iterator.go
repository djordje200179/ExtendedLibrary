package collections

type Iterator[T any] interface {
	Valid() bool
	Move()

	Get() T
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

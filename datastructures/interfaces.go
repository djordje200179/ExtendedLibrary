package datastructures

type Iterator[T any] interface {
	IsValid() bool
	Move()

	Get() T
}

type Sizer interface {
	Size() int
}

type Indexer[K, V any] interface {
	Get(K) V
}

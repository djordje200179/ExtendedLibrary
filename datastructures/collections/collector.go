package collections

type Collector[T any] struct {
	Collection Collection[T]
}

func (collector Collector[T]) Supply(value T) {
	collector.Collection.Append(value)
}

func (collector Collector[T]) Finish() Collection[T] {
	return collector.Collection
}

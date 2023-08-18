package sets

type Collector[T any] struct {
	Set Set[T]
}

func (collector Collector[T]) Supply(value T) {
	collector.Set.Add(value)
}

func (collector Collector[T]) Finish() Set[T] {
	return collector.Set
}

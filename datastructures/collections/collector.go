package collections

type Collector[T any, C Collection[T]] struct {
	Collection C
}

func (collector Collector[T, C]) Supply(value T) {
	collector.Collection.Append(value)
}

func (collector Collector[T, C]) Finish() C {
	return collector.Collection
}

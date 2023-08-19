package sets

type Collector[T any, S Set[T]] struct {
	Set S
}

func (collector Collector[T, S]) Supply(value T) {
	collector.Set.Add(value)
}

func (collector Collector[T, S]) Finish() S {
	return collector.Set
}

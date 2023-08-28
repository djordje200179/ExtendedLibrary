package cols

// Collector is a streams.Collector that can be used to collect values into a Collection.
type Collector[T any, C Collection[T]] struct {
	Collection C // The collection to collect values into.
}

// Supply appends the value to the end of the Collection.
func (collector Collector[T, C]) Supply(value T) {
	collector.Collection.Append(value)
}

// Finish returns the Collection.
func (collector Collector[T, C]) Finish() C {
	return collector.Collection
}

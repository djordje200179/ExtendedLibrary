package cols

// Collector is a generic type that can be used to collect values into a collection.
// It is used by the stream to collect values into a collection.
type Collector[T any, C Collection[T]] struct {
	Collection C // The collection to collect values into.
}

// Supply adds a value to the end of the collection.
func (collector Collector[T, C]) Supply(value T) {
	collector.Collection.Append(value)
}

// Finish returns the collection.
func (collector Collector[T, C]) Finish() C {
	return collector.Collection
}

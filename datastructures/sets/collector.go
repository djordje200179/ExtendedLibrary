package sets

// Collector is a generic type that can be used to collect values into a set.
// It is used by the stream to collect values into a set.
type Collector[T any, S Set[T]] struct {
	Set S // The set to collect values into.
}

// Supply adds a value to the set.
func (collector Collector[T, S]) Supply(value T) {
	collector.Set.Add(value)
}

// Finish returns the set.
func (collector Collector[T, S]) Finish() S {
	return collector.Set
}

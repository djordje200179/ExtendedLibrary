package seqs

// BackPusher is a type that can push a value to the back of a sequence.
type BackPusher[T any] interface {
	PushBack(value T) // Pushes a value to the back of a sequence.
}

// Collector is a type that can collect values and push them to the back of a sequence.
type Collector[T any, BP BackPusher[T]] struct {
	BackPusher BP
}

// Supply pushes a value to the back of the sequence.
func (collector Collector[T, BP]) Supply(value T) {
	collector.BackPusher.PushBack(value)
}

// Finish returns the BackPusher.
func (collector Collector[T, BP]) Finish() BP {
	return collector.BackPusher
}

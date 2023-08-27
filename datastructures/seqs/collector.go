package seqs

type BackPusher[T any] interface {
	PushBack(value T)
}

type Collector[T any, BP BackPusher[T]] struct {
	BackPusher BP
}

func (collector Collector[T, BP]) Supply(value T) {
	collector.BackPusher.PushBack(value)
}

func (collector Collector[T, BP]) Finish() BP {
	return collector.BackPusher
}

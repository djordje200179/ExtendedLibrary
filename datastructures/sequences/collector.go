package sequences

type BackPusher[T any] interface {
	PushBack(value T)
}

type Collector[T any, Sequence BackPusher[T]] struct {
	BackPusher[T]
}

func (collector Collector[T, Sequence]) Supply(value T) {
	collector.BackPusher.PushBack(value)
}

func (collector Collector[T, Sequence]) Finish() Sequence {
	return collector.BackPusher.(Sequence)
}

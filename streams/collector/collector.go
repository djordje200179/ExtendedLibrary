package collector

type Collector[T, R any] interface {
	Supply(value T)
	Finish() R
}

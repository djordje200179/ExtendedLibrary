package mapreduce

type Emitter[K comparable, V any] func(key K, value V)

type Mapper[K comparable, V any] interface {
	Map(emit Emitter[K, V])
}

type Reducer[K comparable, V any] func(key K, values []V) V
type Finalizer[K comparable, V any] func(key K, value V) V

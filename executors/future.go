package executors

import "sync"

type Future[T any] struct {
	getter func() (T, error)

	result T
	err    error

	wg sync.WaitGroup
}

func NewFuture[T any](getter func() (T, error)) *Future[T] {
	future := &Future[T]{
		getter: getter,
	}

	future.wg.Add(1)

	return future
}

func NewErrorlessFuture[T any](getter func() T) *Future[T] {
	return NewFuture[T](func() (T, error) {
		return getter(), nil
	})
}

func (future *Future[T]) TryGet() (T, error) {
	future.wg.Wait()

	return future.result, future.err
}

func (future *Future[T]) Get() T {
	future.wg.Wait()

	if future.err != nil {
		panic(future.err)
	}

	return future.result
}

func (future *Future[T]) Task() Task {
	return func() {
		future.result, future.err = future.getter()
		future.wg.Done()
	}
}

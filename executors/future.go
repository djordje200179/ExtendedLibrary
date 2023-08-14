package executors

import "sync"

type Future[T any] struct {
	getter func() (T, error)
	result T
	err    error

	wg      sync.WaitGroup
	started bool
	done    bool
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

func (future *Future[T]) Function() func() {
	return func() {
		future.result, future.err = future.getter()
	}
}

func (future *Future[T]) IsStarted() bool {
	return future.started
}

func (future *Future[T]) IsDone() bool {
	return future.done
}

func (future *Future[T]) MarkStarted() {
	future.started = true
}

func (future *Future[T]) MarkDone() {
	future.done = true
	future.wg.Done()
}

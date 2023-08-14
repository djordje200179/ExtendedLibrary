package executors

type Future[T any] struct {
	getter func() (T, error)
	result T
	err    error

	*taskWaiter
}

func NewFuture[T any](getter func() (T, error)) *Future[T] {
	future := &Future[T]{
		getter:     getter,
		taskWaiter: newTaskWaiter(),
	}

	return future
}

func NewErrorlessFuture[T any](getter func() T) *Future[T] {
	return NewFuture[T](func() (T, error) {
		return getter(), nil
	})
}

func (future *Future[T]) TryGet() (T, error) {
	future.Wait()

	return future.result, future.err
}

func (future *Future[T]) Get() T {
	future.Wait()

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

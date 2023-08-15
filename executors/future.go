package executors

import "context"

type Future[T any] struct {
	getter func(context.Context) T
	result T

	*taskStatus
}

func NewFuture[T any](getter func(context.Context) T, ctx context.Context) *Future[T] {
	future := &Future[T]{
		getter:     getter,
		taskStatus: newTaskStatus(ctx),
	}

	return future
}

func NewDefaultFuture[T any](getter func(context.Context) T) *Future[T] {
	return NewFuture(getter, context.Background())
}

func (future *Future[T]) Result() T {
	future.Wait()

	if future.IsFailed() {
		panic(future.taskStatus.failReason)
	}

	return future.result
}

func (future *Future[T]) Function() func(context.Context) {
	return func(ctx context.Context) {
		future.result = future.getter(ctx)
	}
}

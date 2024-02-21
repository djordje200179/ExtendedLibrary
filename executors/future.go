package executors

import "context"

type Future[T any] struct {
	getter func(context.Context) T
	result T

	*status
}

func NewFuture[T any](getter func(context.Context) T, ctx context.Context) *Future[T] {
	future := &Future[T]{
		getter: getter,
		status: newTaskStatus(ctx),
	}

	return future
}

func NewDefaultFuture[T any](getter func(context.Context) T) *Future[T] {
	return NewFuture(getter, context.Background())
}

func (f *Future[T]) Result() T {
	f.Wait()

	if f.HasFailed() {
		panic(f.status.failReason)
	}

	return f.result
}

func (f *Future[T]) Function() func(context.Context) {
	return func(ctx context.Context) {
		f.result = f.getter(ctx)
	}
}

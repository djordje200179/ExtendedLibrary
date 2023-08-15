package executors

import "context"

type Task interface {
	Function() func(ctx context.Context)
	Context() context.Context

	IsStarted() bool
	IsFailed() bool
	IsCancelled() bool
	IsFinished() bool

	Wait()
	Cancel()

	MarkStarted()
	MarkFailed(reason any)
	MarkFinished()
}

package executors

import "context"

type Task interface {
	Function() func(ctx context.Context)
	Context() context.Context

	HasStarted() bool
	HasFailed() bool
	HasCancelled() bool
	HasFinished() bool

	Wait()
	Cancel()

	MarkStart()
	MarkFail(reason any)
	MarkFinish()
}

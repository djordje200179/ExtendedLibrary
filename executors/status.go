package executors

import (
	"context"
	"sync"
)

type status struct {
	wg sync.WaitGroup

	context context.Context
	cancel  context.CancelFunc

	failReason any

	started, finished bool
	failed, canceled  bool
}

func newTaskStatus(ctx context.Context) *status {
	taskContext, taskCancel := context.WithCancel(ctx)

	s := &status{
		context: taskContext,
		cancel:  taskCancel,
	}

	s.wg.Add(1)

	return s
}

func (status *status) Context() context.Context {
	return status.context
}

func (status *status) Wait() { status.wg.Wait() }

func (status *status) Cancel() {
	status.canceled = true
	status.cancel()
	status.wg.Done()
}

func (status *status) HasStarted() bool {
	return status.started
}
func (status *status) HasFailed() bool {
	return status.failed
}
func (status *status) HasCancelled() bool {
	return status.canceled
}
func (status *status) HasFinished() bool {
	return status.finished
}

func (status *status) MarkStart() { status.started = true }

func (status *status) MarkFail(failReason any) {
	status.failed = true
	status.failReason = failReason
}

func (status *status) MarkFinish() {
	status.finished = true
	status.wg.Done()
}

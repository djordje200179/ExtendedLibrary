package executors

import (
	"context"
	"sync"
)

type taskStatus struct {
	wg sync.WaitGroup

	context context.Context
	cancel  context.CancelFunc

	failReason any

	started, finished bool
	failed, canceled  bool
}

func newTaskStatus(ctx context.Context) *taskStatus {
	taskContext, taskCancel := context.WithCancel(ctx)

	status := &taskStatus{
		context: taskContext,
		cancel:  taskCancel,
	}

	status.wg.Add(1)

	return status
}

func (status *taskStatus) Context() context.Context {
	return status.context
}

func (status *taskStatus) Wait() {
	status.wg.Wait()
}

func (status *taskStatus) Cancel() {
	status.canceled = true
	status.cancel()
	status.wg.Done()
}

func (status *taskStatus) IsStarted() bool {
	return status.started
}

func (status *taskStatus) IsFailed() bool {
	return status.failed
}

func (status *taskStatus) IsCancelled() bool {
	return status.canceled
}

func (status *taskStatus) IsFinished() bool {
	return status.finished
}

func (status *taskStatus) MarkStarted() {
	status.started = true
}

func (status *taskStatus) MarkFailed(failReason any) {
	status.failed = true
	status.failReason = failReason
}

func (status *taskStatus) MarkFinished() {
	status.finished = true
	status.wg.Done()
}

package executors

import "sync"

type taskWaiter struct {
	wg sync.WaitGroup

	started bool
	done    bool
}

func newTaskWaiter() *taskWaiter {
	waiter := &taskWaiter{}

	waiter.wg.Add(1)

	return waiter
}

func (waiter *taskWaiter) Wait() {
	waiter.wg.Wait()
}

func (waiter *taskWaiter) IsStarted() bool {
	return waiter.started
}

func (waiter *taskWaiter) IsDone() bool {
	return waiter.done
}

func (waiter *taskWaiter) MarkStarted() {
	waiter.started = true
}

func (waiter *taskWaiter) MarkDone() {
	waiter.done = true
	waiter.wg.Done()
}

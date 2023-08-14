package executors

import "sync"

type Action struct {
	function func()

	wg      sync.WaitGroup
	started bool
	done    bool
}

func NewAction(function func()) *Action {
	action := &Action{
		function: function,
	}

	action.wg.Add(1)

	return action
}

func (action *Action) Function() func() {
	return action.function
}

func (action *Action) Wait() {
	action.wg.Wait()
}

func (action *Action) IsStarted() bool {
	return action.started
}

func (action *Action) IsDone() bool {
	return action.done
}

func (action *Action) MarkStarted() {
	action.started = true
}

func (action *Action) MarkDone() {
	action.done = true
	action.wg.Done()
}

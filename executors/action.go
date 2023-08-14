package executors

type Action struct {
	function func()

	*taskWaiter
}

func NewAction(function func()) *Action {
	action := &Action{
		function:   function,
		taskWaiter: newTaskWaiter(),
	}

	return action
}

func (action *Action) Function() func() {
	return action.function
}

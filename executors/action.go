package executors

import "context"

type Action struct {
	function func(context.Context)

	*taskStatus
}

func NewAction(function func(context.Context), ctx context.Context) *Action {
	action := &Action{
		function:   function,
		taskStatus: newTaskStatus(ctx),
	}

	return action
}

func NewDefaultAction(function func(context.Context)) *Action {
	return NewAction(function, context.Background())
}

func (action *Action) Function() func(context.Context) {
	return action.function
}

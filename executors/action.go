package executors

import "context"

type Action struct {
	f func(context.Context)

	*status
}

func NewAction(f func(context.Context), ctx context.Context) *Action {
	action := &Action{
		f: f,

		status: newTaskStatus(ctx),
	}

	return action
}

func NewDefaultAction(function func(context.Context)) *Action {
	return NewAction(function, context.Background())
}

func (a *Action) Function() func(context.Context) {
	return a.f
}

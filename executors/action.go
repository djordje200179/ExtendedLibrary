package executors

type Action func()

func NewAction(action func()) Action {
	return action
}

func (action Action) Task() Task {
	return Task(action)
}

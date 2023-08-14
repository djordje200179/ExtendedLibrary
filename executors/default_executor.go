package executors

import "runtime"

var DefaultExecutor = NewExecutor(runtime.NumCPU())

func Submit(tasker Tasker) {
	DefaultExecutor.Submit(tasker)
}

package executors

import "runtime"

var DefaultExecutor = NewExecutor(runtime.NumCPU())

func Submit(task Task) {
	DefaultExecutor.Submit(task)
}

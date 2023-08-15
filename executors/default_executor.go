package executors

import "runtime"

var DefaultExecutor = NewExecutor(runtime.NumCPU(), 100)

func Submit(task Task) {
	DefaultExecutor.Submit(task)
}

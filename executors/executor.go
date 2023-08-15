package executors

import (
	"sync"
)

type Executor struct {
	tasks chan Task

	wg sync.WaitGroup
}

func NewExecutor(goroutinesCount int, queueSize int) *Executor {
	executor := &Executor{
		tasks: make(chan Task, queueSize),
	}

	for i := 0; i < goroutinesCount; i++ {
		go executor.routine()
	}

	executor.wg.Add(goroutinesCount)

	return executor
}

func (executor *Executor) Submit(tasker Task) {
	executor.tasks <- tasker
}

func (executor *Executor) Close() {
	close(executor.tasks)
	executor.wg.Wait()
}

func (executor *Executor) routine() {
	for task := range executor.tasks {
		runTask(task)
	}

	executor.wg.Done()
}

func runTask(task Task) {
	defer func() {
		if r := recover(); r != nil {
			task.MarkFailed(r)
		}
	}()

	function := task.Function()
	context := task.Context()

	task.MarkStarted()
	function(context)
	task.MarkFinished()
}

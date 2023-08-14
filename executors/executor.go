package executors

import (
	"sync"
)

type Executor struct {
	tasks chan Task

	wg sync.WaitGroup
}

func NewExecutor(goroutinesCount int) *Executor {
	executor := &Executor{
		tasks: make(chan Task, goroutinesCount),
	}

	for i := 0; i < goroutinesCount; i++ {
		go executor.routine()
	}

	executor.wg.Add(goroutinesCount)

	return executor
}

func (executor *Executor) Submit(tasker Tasker) {
	task := tasker.Task()
	executor.tasks <- task
}

func (executor *Executor) Close() {
	close(executor.tasks)
	executor.wg.Wait()
}

func (executor *Executor) routine() {
	for task := range executor.tasks {
		task()
	}
	executor.wg.Done()
}

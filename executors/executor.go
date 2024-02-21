package executors

import (
	"sync"
)

type Executor struct {
	tasks chan Task

	wg sync.WaitGroup
}

func NewExecutor(goroutinesCnt int, queueSize int) *Executor {
	executor := &Executor{
		tasks: make(chan Task, queueSize),
	}

	for range goroutinesCnt {
		go executor.routine()
	}

	executor.wg.Add(goroutinesCnt)

	return executor
}

func (e *Executor) Submit(tasker Task) { e.tasks <- tasker }

func (e *Executor) Close() {
	close(e.tasks)
	e.wg.Wait()
}

func (e *Executor) routine() {
	for task := range e.tasks {
		runTask(task)
	}

	e.wg.Done()
}

func runTask(task Task) {
	defer func() {
		if r := recover(); r != nil {
			task.MarkFail(r)
		}
	}()

	function := task.Function()
	context := task.Context()

	task.MarkStart()
	function(context)
	task.MarkFinish()
}

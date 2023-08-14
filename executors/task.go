package executors

type Task func()

type Tasker interface {
	Task() Task
}

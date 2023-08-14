package executors

type Task interface {
	Function() func()

	IsStarted() bool
	IsDone() bool

	MarkStarted()
	MarkDone()
}

package periodictask

import "time"

type signal bool

const (
	pause  signal = true
	resume signal = false
)

type PeriodicTask struct {
	task   func()
	period time.Duration
	async  bool

	signal  chan signal
	stopped bool
}

func New(task func(), period time.Duration, async bool) *PeriodicTask {
	pTask := new(PeriodicTask)

	pTask.task = task
	pTask.period = period
	pTask.async = async
	pTask.signal = make(chan signal, 1)

	return pTask
}

func (task *PeriodicTask) wrapper() {
	for !task.stopped {
		select {
		case signal := <-task.signal:
			if signal == pause {
				for <-task.signal == pause {
				}
			}
		default:
			if task.async {
				go task.task()
			} else {
				task.task()
			}

			time.Sleep(task.period)
		}
	}
}

func (task *PeriodicTask) Start()  { go task.wrapper() }
func (task *PeriodicTask) Stop()   { task.stopped = true }
func (task *PeriodicTask) Pause()  { task.signal <- pause }
func (task *PeriodicTask) Resume() { task.signal <- resume }

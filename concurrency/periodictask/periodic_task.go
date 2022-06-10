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
	return &PeriodicTask{
		task:   task,
		period: period,
		async:  async,
		signal: make(chan signal),
	}
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

package event

import "github.com/djordje200179/extendedlibrary/concurrency/semaphore"

type Event struct {
	sem semaphore.Semaphore
}

func New() Event {
	return Event{semaphore.New(0)}
}

func (event *Event) Wait() {
	event.sem.Wait()
}

func (event *Event) Notify() {
	event.sem.Signal()
}

func (event *Event) NotifyAll() {
	for event.sem.Value() > 0 {
		event.sem.Signal()
	}
}

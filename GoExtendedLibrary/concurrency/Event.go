package concurrency

import "github.com/djordje200179/GoExtendedLibrary/concurrency/semaphore"

type Event semaphore.Semaphore

func New() Event {
	return Event(semaphore.New(0))
}

func (event *Event) getSem() *semaphore.Semaphore {
	return (*semaphore.Semaphore)(event)
}

func (event *Event) Wait() {
	event.getSem().Wait()
}

func (event *Event) Notify() {
	event.getSem().Signal()
}

func (event *Event) NotifyAll() {
	for event.getSem().Value() > 0 {
		event.getSem().Signal()
	}
}

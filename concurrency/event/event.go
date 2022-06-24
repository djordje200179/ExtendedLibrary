package event

import (
	"sync"
)

type Event struct {
	cond *sync.Cond
}

func New() Event { return Event{sync.NewCond(&sync.Mutex{})} }

func (event Event) Wait() {
	event.cond.L.Lock()
	event.cond.Wait()
	event.cond.L.Unlock()
}

func (event Event) Notify()    { event.cond.Signal() }
func (event Event) NotifyAll() { event.cond.Broadcast() }

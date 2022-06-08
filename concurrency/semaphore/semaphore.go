package semaphore

type Semaphore struct {
}

func New(val int) Semaphore {
	return Semaphore{}
}

func (s *Semaphore) Wait() {

}

func (s *Semaphore) Signal() {

}

func (s *Semaphore) Value() int {
	return 0
}

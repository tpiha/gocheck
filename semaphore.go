package main

type Semaphore chan bool

// Lock locks the semaphore until unlocked, passes until concurrecny is reached and then blocks
func (s Semaphore) Lock() {
	<-s
}

// Unlock unlocks the semaphore
func (s Semaphore) Unlock() {
	s <- true
}

// NewSemaphore creates a new semaphore object
func NewSemaphore(concurrency uint) Semaphore {
	s := make(Semaphore, concurrency)

	var i uint

	for i = 0; i < concurrency; i++ {
		s.Unlock()
	}

	return s
}

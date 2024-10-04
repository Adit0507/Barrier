package barrier

import "sync"

type Barrier struct {
	size      int // no. of gr. in barrier
	waitCount int // no. of currently suspended executions
	cond      *sync.Cond
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount += 1

	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}

	b.cond.L.Unlock()
}

func NewBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})
	return &Barrier{size, 0, condVar}
}

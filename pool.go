package async

import (
	"sync"

	"github.com/y00rb/async/scheduler"
	"github.com/y00rb/async/worker"
)

type Pool struct {
	scheduler   scheduler.Scheduler
	workerCount int
	quit        chan struct{}
	wg          sync.WaitGroup
}

func NewPool(size int) *Pool {
	ce := Pool{
		scheduler:   &scheduler.FuncSchedule{},
		workerCount: size,
		quit:        make(chan struct{}, size),
	}
	ce.run()
	return &ce
}

func (p *Pool) Submit(executor worker.Executor) {
	p.wg.Add(1)
	p.scheduler.Submit(executor)
}

func (ce *Pool) run() {
	ce.scheduler.Run()
	for i := 0; i < ce.workerCount; i++ {
		ce.createWorker()
	}
}

func (ce *Pool) Quit() {
	ce.wg.Wait()
	ce.scheduler.Stop()
	for i := 0; i < ce.workerCount; i++ {
		ce.quit <- struct{}{}
	}

}

func (ce *Pool) createWorker() {
	go func(in chan worker.Executor, ready worker.ReadyResponse) {
		for {
			ready.WorkerReady(in)
			select {
			case execcutor := <-in:
				execcutor.Exec()
				ce.wg.Done()
			case <-ce.quit:
				return
			}
		}
	}(ce.scheduler.WorkerChan(), ce.scheduler)
}

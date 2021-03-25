package async

import (
	"github.com/y00rb/async/scheduler"
	"github.com/y00rb/async/worker"
)

type Pool struct {
	scheduler   scheduler.Scheduler
	workerCount int
}

func NewPool(size int) *Pool {
	ce := Pool{
		scheduler:   &scheduler.FuncSchedule{},
		workerCount: size,
	}
	ce.run()
	return &ce
}

func (p *Pool) Submit(executor worker.Executor) {
	p.scheduler.Submit(executor)
}

func (ce *Pool) run() {
	ce.scheduler.Run()
	for i := 0; i < ce.workerCount; i++ {
		createWorker(ce.scheduler.WorkerChan(), ce.scheduler)
	}
}

func createWorker(in chan worker.Executor, ready worker.ReadyResponse) {
	go func(in chan worker.Executor) {
		for {
			ready.WorkerReady(in)
			executor := <-in

			err := workerExec(executor)

			if err != nil {
				// TODO: catch the error
				continue
			}
		}
	}(in)
}

func workerExec(r worker.Executor) error {
	r.Exec()
	return nil
}

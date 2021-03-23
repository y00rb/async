package async

import (
	"github.com/y00rb/async/scheduler"
	"github.com/y00rb/async/worker"
)

type Pool struct {
	scheduler   scheduler.Scheduler
	workerCount int
	poolFunc    func(interface{})
}

func NewPool(size int) *Pool {
	ce := Pool{
		scheduler:   &scheduler.FuncScheduler{},
		workerCount: size,
	}
	ce.run()
	return &ce
}

func (p *Pool) Submit(req worker.Request) {
	p.scheduler.Submit(req)
}

func (ce *Pool) run() {
	ce.scheduler.Run()
	for i := 0; i < ce.workerCount; i++ {
		createWorker(ce.scheduler.WorkerChan(), ce.scheduler)
	}
}

func createWorker(in chan worker.Request, ready worker.ReadyResponse) {
	go func(in chan worker.Request) {
		for {
			ready.WorkerReady(in)
			request := <-in

			err := workerExec(request)

			if err != nil {
				// TODO: catch the error
				continue
			}
		}
	}(in)
}

func workerExec(r worker.Request) error {
	r()
	return nil
}

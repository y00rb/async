package async

import (
	"github.com/y00rb/async/scheduler"
	"github.com/y00rb/async/worker"
)

type Pool struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
	poolFunc    func(interface{})
}

func NewPool(size int) *Pool {
	ce := Pool{
		Scheduler:   &scheduler.FuncScheduler{},
		WorkerCount: size,
	}
	ce.run()
	return &ce
}

func (ce *Pool) run() {
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChan(), ce.Scheduler)
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

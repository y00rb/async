package async

import (
	"github.com/y00rb/async/scheduler"
	"github.com/y00rb/async/worker"
)

type Pool struct {
	scheduler   scheduler.Scheduler
	workerCount int
	workerFunc  worker.Func
}

func NewPool(size int, wf worker.Func) *Pool {
	ce := Pool{
		scheduler:   &scheduler.FuncSchedule{},
		workerCount: size,
		workerFunc:  wf,
	}
	ce.run()
	return &ce
}

func (p *Pool) Submit(params worker.Params) {
	p.scheduler.Submit(params)
}

func (ce *Pool) run() {
	ce.scheduler.Run()
	for i := 0; i < ce.workerCount; i++ {
		createWorker(ce.workerFunc, ce.scheduler.WorkerChan(), ce.scheduler)
	}
}

func createWorker(wf worker.Func, in chan worker.Params, ready worker.ReadyResponse) {
	go func(wf worker.Func, in chan worker.Params) {
		for {
			ready.WorkerReady(in)
			params := <-in

			err := workerExec(wf, params)

			if err != nil {
				// TODO: catch the error
				continue
			}
		}
	}(wf, in)
}

func workerExec(wf worker.Func, r worker.Params) error {
	wf(r)
	return nil
}

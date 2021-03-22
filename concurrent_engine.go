package async

import (
	"context"

	"github.com/y00rb/async/engine"
	"github.com/y00rb/async/scheduler"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
}

func (ce *ConcurrentEngine) Run(ctx context.Context) {
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChan(), ce.Scheduler)
	}
}

func createWorker(in chan engine.Request, ready engine.ReadyResponse) {
	go func(in chan engine.Request) {
		for {
			ready.WorkerReady(in)
			request := <-in

			err := worker(request)

			if err != nil {
				// TODO: catch the error
				continue
			}
		}
	}(in)
}

func worker(r engine.Request) error {
	r()
	return nil
}

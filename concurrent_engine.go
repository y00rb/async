package dragonfly

import (
	"context"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (ce *ConcurrentEngine) Run(ctx context.Context) {
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChan(), ce.Scheduler)
	}
}

func createWorker(in chan Request, ready ReadyResponse) {
	go func(in chan Request) {
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

func worker(r Request) error {
	r()
	return nil
}

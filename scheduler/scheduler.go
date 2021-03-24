package scheduler

import "github.com/y00rb/async/worker"

type Scheduler interface {
	worker.ReadyResponse
	Submit(worker.Params)
	WorkerChan() chan worker.Params
	Run()
}

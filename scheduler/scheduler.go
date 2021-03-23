package scheduler

import "github.com/y00rb/async/worker"

type Scheduler interface {
	worker.ReadyResponse
	Submit(worker.Request)
	WorkerChan() chan worker.Request
	Run()
}

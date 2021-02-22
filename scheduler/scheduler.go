package scheduler

import "github.com/y00rb/dragonfly/engine"

type Scheduler interface {
	engine.ReadyResponse
	Submit(engine.Request)
	WorkerChan() chan engine.Request
	Run()
}

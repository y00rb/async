package engine

type Request func()

type ReadyResponse interface {
	WorkerReady(chan Request)
}
package dragonfly

type Scheduler interface {
	ReadyResponse
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

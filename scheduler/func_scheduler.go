package scheduler

import "github.com/y00rb/async/engine"

type FuncScheduler struct {
	request    chan engine.Request
	workerChan chan chan engine.Request
}

func (f *FuncScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (f *FuncScheduler) Submit(r engine.Request) {
	f.request <- r
}

func (f *FuncScheduler) WorkerReady(w chan engine.Request) {
	f.workerChan <- w
}

func (f *FuncScheduler) Run() {
	f.request = make(chan engine.Request)
	f.workerChan = make(chan chan engine.Request)

	go func() {
		var (
			requestQ []engine.Request
			workerQ  []chan engine.Request
		)
		for {
			var (
				activeRequest engine.Request
				activeWorker  chan engine.Request
			)
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-f.request:
				requestQ = append(requestQ, r)
			case w := <-f.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

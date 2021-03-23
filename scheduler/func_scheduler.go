package scheduler

import "github.com/y00rb/async/worker"

type FuncScheduler struct {
	request    chan worker.Request
	workerChan chan chan worker.Request
}

func (f *FuncScheduler) WorkerChan() chan worker.Request {
	return make(chan worker.Request)
}

func (f *FuncScheduler) Submit(r worker.Request) {
	f.request <- r
}

func (f *FuncScheduler) WorkerReady(w chan worker.Request) {
	f.workerChan <- w
}

func (f *FuncScheduler) Run() {
	f.request = make(chan worker.Request)
	f.workerChan = make(chan chan worker.Request)

	go func() {
		var (
			requestQ []worker.Request
			workerQ  []chan worker.Request
		)
		for {
			var (
				activeRequest worker.Request
				activeWorker  chan worker.Request
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

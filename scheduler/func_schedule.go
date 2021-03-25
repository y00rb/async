package scheduler

import "github.com/y00rb/async/worker"

type FuncSchedule struct {
	params     chan worker.Params
	workerChan chan chan worker.Params
}

func (f *FuncSchedule) WorkerChan() chan worker.Params {
	return make(chan worker.Params)
}

func (f *FuncSchedule) Submit(r worker.Params) {
	f.params <- r
}

func (f *FuncSchedule) WorkerReady(w chan worker.Params) {
	f.workerChan <- w
}

func (f *FuncSchedule) Run() {
	f.params = make(chan worker.Params)
	f.workerChan = make(chan chan worker.Params)

	go func() {
		var (
			requestQ []worker.Params
			workerQ  []chan worker.Params
		)
		for {
			var (
				activeRequest worker.Params
				activeWorker  chan worker.Params
			)
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-f.params:
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

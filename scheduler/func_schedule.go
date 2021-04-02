package scheduler

import (
	"github.com/y00rb/async/worker"
)

type FuncSchedule struct {
	executor   chan worker.Executor
	workerChan chan chan worker.Executor
	stop       chan struct{}
}

func (f *FuncSchedule) WorkerChan() chan worker.Executor {
	return make(chan worker.Executor)
}

func (f *FuncSchedule) Submit(r worker.Executor) {
	f.executor <- r
}

func (f *FuncSchedule) WorkerReady(w chan worker.Executor) {
	f.workerChan <- w
}

func (f *FuncSchedule) Run() {
	f.executor = make(chan worker.Executor)
	f.workerChan = make(chan chan worker.Executor)
	f.stop = make(chan struct{})

	go func() {
		var (
			requestQ []worker.Executor
			workerQ  []chan worker.Executor
		)
		for {
			var (
				activeRequest worker.Executor
				activeWorker  chan worker.Executor
			)
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-f.executor:
				requestQ = append(requestQ, r)
			case w := <-f.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			case <-f.stop:
				return
			}
		}
	}()
}

func (f *FuncSchedule) Stop() {
	f.stop <- struct{}{}
}

package dragonfly

type FuncScheduler struct {
	request    chan Request
	workerChan chan chan Request
}

func (f *FuncScheduler) WorkerChan() chan Request {
	return make(chan Request)
}

func (f *FuncScheduler) Submit(r Request) {
	f.request <- r
}

func (f *FuncScheduler) WorkerReady(w chan Request) {
	f.workerChan <- w
}

func (f *FuncScheduler) Run() {
	f.request = make(chan Request)
	f.workerChan = make(chan chan Request)

	go func() {
		var (
			requestQ []Request
			workerQ  []chan Request
		)
		for {
			var (
				activeRequest Request
				activeWorker  chan Request
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

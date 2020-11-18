package dragonfly

import "fmt"

type Engine struct {
	WorkerCount int
	workerChan  chan struct{}
}

func (e *Engine) Run(f func()) {
	e.workerChan <- struct{}{}
	go func() {
		fmt.Println("executing")
		f()
		<-e.workerChan
		fmt.Println("executed")
		fmt.Println("end")
	}()
}

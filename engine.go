package dragonfly

import "fmt"

type Engine struct {
	WorkerCount int
	workerChan  chan struct{}
}

func NewEngine(count int) *Engine {
	return &Engine{
		WorkerCount: count,
		workerChan:  make(chan struct{}, count),
	}
}

func (e *Engine) ExecTask(f func()) {
	e.workerChan <- struct{}{}
	go func() {
		fmt.Println("executing")
		f()
		<-e.workerChan
		fmt.Println("executed")
		fmt.Println("end")
	}()
}

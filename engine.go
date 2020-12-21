package dragonfly

import (
	"fmt"
)

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

func (e *Engine) Run(f func()) {
	e.workerChan <- struct{}{}
	go func() {
		// fmt.Println("executing")
		f()
		<-e.workerChan
		// fmt.Println("executed")
		// fmt.Println("end")
	}()
}

func (e *Engine) RunWithArgs(f func(t interface{}), p interface{}) {
	e.workerChan <- struct{}{}
	go func(p interface{}) {
		fmt.Println("executing")
		f(p)
		<-e.workerChan
		fmt.Println("executed")
		fmt.Println("end")
	}(p)
}

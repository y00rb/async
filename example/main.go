package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/y00rb/dragonfly"
	"github.com/y00rb/dragonfly/scheduler"
)

var runTimes = 100

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	ce := dragonfly.ConcurrentEngine{
		Scheduler:   &scheduler.FuncScheduler{},
		WorkerCount: 10,
	}

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	ce.Run(ctx)
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			ce.Scheduler.Submit(syncCalculateSum)
		}
	}()
	wg.Wait()
	cancelFunc()
}

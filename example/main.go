package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/y00rb/dragonfly"
)

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	ce := dragonfly.ConcurrentEngine{
		Scheduler:   &dragonfly.FuncScheduler{},
		WorkerCount: 10,
	}

	var wg sync.WaitGroup
	syncCalculateSum := func() {

		demoFunc()
		wg.Done()
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			ce.Scheduler.Submit(syncCalculateSum)
		}
		cancelFunc()
	}()
	wg.Wait()
	ce.Run(ctx)
}

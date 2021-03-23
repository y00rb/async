package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/y00rb/async"
)

var runTimes = 100

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	ce := async.NewPool(10)

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			ce.Scheduler.Submit(syncCalculateSum)
		}
	}()
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/y00rb/async"
)

var runTimes = 100

func main() {
	var wg sync.WaitGroup
	syncCalculateSum := func(i interface{}) {
		count := i.(int)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!", count)
		wg.Done()
	}
	ce := async.NewPool(10, syncCalculateSum)
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			ce.Submit(i)
		}
	}()
	wg.Wait()
}

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
	demoFuncWithParams := func(i interface{}) {
		count := i.(int)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!", count)
		wg.Done()
	}
	ce := async.NewPool(10, demoFuncWithParams)
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			ce.Submit(i)
		}
	}()
	wg.Wait()

	// demoFunc := func() {
	// 	time.Sleep(10 * time.Millisecond)
	// 	fmt.Println("Hello World!")
	// 	wg.Done()
	// }
	// pool := async.NewPool(10, demoFunc)
	// wg.Add(runTimes)
	// go func() {
	// 	for i := 0; i < runTimes; i++ {
	// 		pool.Submit(i)
	// 	}
	// }()
	// wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/y00rb/async"
	"github.com/y00rb/async/worker"
)

var runTimes = 100

func main() {
	var wg sync.WaitGroup
	demoFuncWithParams := func(i interface{}) {
		count := i.(int)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!", count)
		defer wg.Done()
	}
	ce := async.NewPool(10)
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			e := worker.ExecuteWithParams{
				FuncWithParams: demoFuncWithParams,
				Params:         i,
			}
			ce.Submit(e)
		}
	}()
	wg.Wait()

	demoFunc := func() error {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!")
		defer wg.Done()
		return nil
	}
	pool := async.NewPool(10)
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			e := worker.Execute{
				Func: demoFunc,
			}
			pool.Submit(e)
		}
	}()
	wg.Wait()
}

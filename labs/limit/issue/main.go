package main

import (
	"fmt"
	"sync"
	"time"

	dragonfly "github.com/y00rb/dragonfly"
)

func main() {
	const jobCount = 100
	wg := sync.WaitGroup{}

	e := dragonfly.NewEngine(10)
	for i := 0; i < jobCount; i++ {
		wg.Add(1)
		fmt.Println("test")
		goFunc := func(t interface{}) {
			n := t.(int)
			fmt.Println("job", n)
			time.Sleep(time.Second)
			wg.Done()
		}
		e.RunWithArgs(goFunc, i)
	}
	wg.Wait()
}

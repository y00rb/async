package dragonfly

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestEngineWaitToGetWorker(t *testing.T) {
	wg := sync.WaitGroup{}
	number := 10
	e := Engine{WorkerCount: 3}
	for i := 0; i < number; i++ {
		wg.Add(1)
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}
		e.Run(goFunc)
	}
	wg.Wait()
}

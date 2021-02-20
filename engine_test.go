package dragonfly

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	_        = 1 << (10 * iota)
	MiB      // 1048576
	AntsSize = 3
	TestSize = 10000
	n        = 10
)

func TestEngineRunWaitToGetWorker(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	e := NewEngine(n)
	// wg := sync.WaitGroup{}

	for i := 0; i < TestSize; i++ {
		// wg.Add(1)
		goFunc := func(i interface{}) {
			// time.Sleep(time.Second)
			n := i.(int)
			fmt.Println("executing ", n)
			// wg.Done()
		}
		e.RunWithArgs(goFunc, i)
	}
	// wg.Wait()
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / MiB
	t.Logf("memory usage:%d MB", usedMem)
}

func TestEngineRunWithArgsWaitToGetWorker(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	wg := sync.WaitGroup{}
	e := NewEngine(3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("test")
		goFunc := func(i interface{}) {
			n := i.(int)
			time.Sleep(time.Duration(n) * time.Second)
			wg.Done()
		}
		e.RunWithArgs(goFunc, i)
	}
	wg.Wait()
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / MiB
	t.Logf("memory usage:%d MB", usedMem)
}

package dragonfly

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
)

const (
	_        = 1 << (10 * iota)
	MiB      // 1048576
	AntsSize = 3
	TestSize = 10000
	n        = 10
)

// TestAntsPoolWaitToGetWorker is used to test waiting to get worker.
func TestAntsPoolWaitToGetWorker(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	var wg sync.WaitGroup
	p, _ := ants.NewPool(n)
	defer p.Release()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			fmt.Println("executing ")
			wg.Done()
		})
	}
	wg.Wait()
	t.Logf("pool, running workers number:%d", p.Running())
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / MiB
	t.Logf("memory usage:%d MB", usedMem)
}

func TestAntsPoolPerMalloc(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	var wg sync.WaitGroup
	p, _ := ants.NewPool(3, ants.WithPreAlloc(true))
	defer p.Release()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			fmt.Printf("go func: %d\n", i)
			time.Sleep(time.Second)
			wg.Done()
		})
	}
	wg.Wait()
	t.Logf("pool, running worker number: %d", p.Running())
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / MiB
	t.Logf("memory usage:%d MB", usedMem)
}

func TestAntsPool(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	defer ants.Release()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = ants.Submit(func() {
			fmt.Printf("go func: %d\n", i)
			time.Sleep(time.Second)
			wg.Done()
		})
	}
	wg.Wait()

	t.Logf("pool, capacity:%d", ants.Cap())
	t.Logf("pool, running workers number:%d", ants.Running())
	t.Logf("pool, free workers number:%d", ants.Free())
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / MiB
	t.Logf("memory usage:%d MB", usedMem)
}

package compare

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/y00rb/async/common"
)

var sumForAnts int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sumForAnts, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func TestAntsPool_WithoutArgs(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)

	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("finish all tasks.\n")

	t.Logf("running goroutines: %d\n", ants.Running())
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / common.MiB
	t.Logf("memory usage:%d MB", usedMem)
}

func TestAntsPool_WithArgs(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)

	defer func() {
		ants.Release()
		sumForAnts = 0
	}()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(20, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()

	fmt.Printf("finish all tasks, result is %d\n", sumForAnts)

	t.Logf("running goroutines: %d\n", p.Running())
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / common.MiB
	t.Logf("memory usage:%d MB", usedMem)
}

package compare

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/y00rb/dragonfly"
	"github.com/y00rb/dragonfly/common"
)

var sum int32 = 0

func myEngineFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoEngineFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func TestEnginePool_WithoutArg(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	e := dragonfly.NewEngine(10)

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoEngineFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		e.Run(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("finish all tasks.\n")

	t.Logf("running goroutines: %d\n", e.WorkerCount)
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / common.MiB
	t.Logf("memory usage:%d MB", usedMem)
}

func TestEnginePool_WithArg(t *testing.T) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)
	e := dragonfly.NewEngine(20)

	runTimes := 1000
	defer func() {
		sum = 0
	}()

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func(i interface{}) {
		myEngineFunc(i)
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		e.RunWithArgs(syncCalculateSum, int32(i))
	}
	wg.Wait()
	fmt.Printf("finish all tasks, result is %d\n", sum)

	t.Logf("running goroutines: %d\n", e.WorkerCount)
	endMem := runtime.MemStats{}
	runtime.ReadMemStats(&endMem)
	usedMem := (endMem.TotalAlloc - startMem.TotalAlloc) / common.MiB
	t.Logf("memory usage:%d MB", usedMem)
}

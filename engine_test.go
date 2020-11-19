package dragonfly

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	_   = 1 << (10 * iota)
	MiB // 1048576
)

var curMem uint64

func TestEngineWaitToGetWorker(t *testing.T) {
	wg := sync.WaitGroup{}
	e := NewEngine(3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("test")
		goFunc := func() {
			time.Sleep(time.Second)
			wg.Done()
		}
		e.ExecTask(goFunc)
	}
	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

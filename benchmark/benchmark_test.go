package benchmark

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/y00rb/async"
	"github.com/y00rb/async/common"
	"github.com/y00rb/async/worker"
)

func BenchmarkGoroutines(b *testing.B) {
	var wg sync.WaitGroup
	count := b.N
	demoFunc := func(i interface{}) {
		count := i.(int)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!", count)
	}
	b.StartTimer()
	for i := 0; i < count; i++ {
		wg.Add(common.RunTimes)
		for j := 0; j < common.RunTimes; j++ {
			go func(j int) {
				demoFunc(j)
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
	b.StopTimer()
	b.Logf("running the demoFunc in %d times without limit\n", common.RunTimes)
}

func BenchmarkConcurrentEngine(b *testing.B) {
	count := common.BenchAntsSize
	var wg sync.WaitGroup
	p := async.NewPool(count)
	defer p.Quit()
	demoFunc := func(i interface{}) {
		count := i.(int)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!", count)
		defer wg.Done()
	}
	wg.Add(common.RunTimes)
	b.StartTimer()
	go func() {
		for i := 0; i < common.RunTimes; i++ {
			executor := worker.ExecuteWithParams{
				FuncWithParams: demoFunc,
				Params:         i,
			}
			p.Submit(executor)
		}
	}()
	wg.Wait()

	b.StopTimer()
	b.Logf("running the demoFunc in %d times in %d goroutines \n", common.RunTimes, count)
}

// func BenchmarkAntsPool(b *testing.B) {
// 	var wg sync.WaitGroup
// 	p, _ := ants.NewPool(common.BenchAntsSize, ants.WithExpiryDuration(common.DefaultExpiredTime))
// 	defer p.Release()

// 	count := b.N
// 	b.StartTimer()
// 	for i := 0; i < count; i++ {
// 		wg.Add(common.RunTimes)
// 		for j := 0; j < common.RunTimes; j++ {
// 			_ = p.Submit(func() {
// 				demoFunc()
// 				wg.Done()
// 			})
// 		}
// 		wg.Wait()
// 		b.Logf("running the demoFunc in %d times in %d goroutines \n", common.RunTimes, common.BenchAntsSize)
// 	}
// 	b.StopTimer()
// }

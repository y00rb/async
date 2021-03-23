package compare

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/y00rb/async"
	"github.com/y00rb/async/common"
)

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func BenchmarkGoroutines(b *testing.B) {
	var wg sync.WaitGroup
	count := b.N
	b.StartTimer()
	for i := 0; i < count; i++ {
		wg.Add(common.RunTimes)
		for j := 0; j < common.RunTimes; j++ {
			go func() {
				demoFunc()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	b.StopTimer()
	b.Logf("running the demoFunc in %d times without limit\n", common.RunTimes)
}

func BenchmarkConcurrentEngine(b *testing.B) {
	count := common.BenchAntsSize
	ce := async.NewPool(count)

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	wg.Add(common.RunTimes)
	b.StartTimer()
	go func() {
		for i := 0; i < common.RunTimes; i++ {
			ce.Scheduler.Submit(syncCalculateSum)
		}
	}()
	wg.Wait()
	b.Logf("running the demoFunc in %d times in %d goroutines \n", common.RunTimes, count)
	b.StopTimer()
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

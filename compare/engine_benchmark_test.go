package compare

import (
	"context"
	"sync"
	"testing"

	"github.com/y00rb/dragonfly"
	"github.com/y00rb/dragonfly/common"

	"github.com/panjf2000/ants/v2"
)

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
	ce := dragonfly.ConcurrentEngine{
		Scheduler:   &dragonfly.FuncScheduler{},
		WorkerCount: count,
	}

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	ce.Run(ctx)
	wg.Add(common.RunTimes)
	b.StartTimer()
	go func() {
		for i := 0; i < common.RunTimes; i++ {
			ce.Scheduler.Submit(syncCalculateSum)
		}
	}()
	wg.Wait()
	b.Logf("running the demoFunc in %d times in %d goroutines \n", common.RunTimes, count)
	cancelFunc()
	b.StopTimer()
}

func BenchmarkAntsPool(b *testing.B) {
	var wg sync.WaitGroup
	p, _ := ants.NewPool(common.BenchAntsSize, ants.WithExpiryDuration(common.DefaultExpiredTime))
	defer p.Release()

	count := b.N
	b.StartTimer()
	for i := 0; i < count; i++ {
		wg.Add(common.RunTimes)
		for j := 0; j < common.RunTimes; j++ {
			_ = p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
		b.Logf("running the demoFunc in %d times in %d goroutines \n", common.RunTimes, common.BenchAntsSize)
	}
	b.StopTimer()
}

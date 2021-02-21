package compare

import (
	"sync"
	"testing"

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
	b.Logf("running the demoFunc in %d times in %d goroutines \n", common.RunTimes, count)
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
		b.Logf("running the demoFunc in %d times in %d goroutines \n", count, common.BenchAntsSize)
	}
	b.StopTimer()
}

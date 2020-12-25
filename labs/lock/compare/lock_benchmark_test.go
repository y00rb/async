package compare

import (
	"sync"
	"testing"
)

func benchmarkLock(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for j := 0; j < write*100; j++ {
			wg.Add(1)
			rw.Write()
			wg.Done()
		}
		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B)    { benchmarkLock(b, &Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B)  { benchmarkLock(b, &RWLock{}, 9, 1) }
func BenchmarkWriteMore(b *testing.B)   { benchmarkLock(b, &Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmarkLock(b, &RWLock{}, 1, 9) }
func BenchmarkEqual(b *testing.B)       { benchmarkLock(b, &Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B)     { benchmarkLock(b, &RWLock{}, 5, 5) }

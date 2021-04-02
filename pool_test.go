package async

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/y00rb/async/worker"
	"go.uber.org/goleak"
)

var runTimes = 100

func TestPool_With10Task(t *testing.T) {
	defer goleak.VerifyNone(t)
	var wg sync.WaitGroup
	demoFuncWithParams := func(i interface{}) {
		count := i.(int)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!", count)
		defer wg.Done()
	}
	p := NewPool(10)
	defer p.Quit()
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			e := worker.ExecuteWithParams{
				FuncWithParams: demoFuncWithParams,
				Params:         i,
			}
			p.Submit(e)
		}
	}()
	wg.Wait()

	demoFunc := func() error {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!")
		defer wg.Done()
		return nil
	}
	pool := NewPool(10)
	defer pool.Quit()
	wg.Add(runTimes)
	go func() {
		for i := 0; i < runTimes; i++ {
			e := worker.Execute{
				Func: demoFunc,
			}
			pool.Submit(e)
		}
	}()
	wg.Wait()
}

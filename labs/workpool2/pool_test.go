package workpool2

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	tasks := []*Task{
		NewTask(func() error { return nil }),
		NewTask(func() error { return nil }),
		NewTask(func() error { return nil }),
	}

	p := NewPool(tasks, 2)
	if len(p.Tasks) == 0 {
		t.Error("tasks should NOT equals 0")
	}
}

func TestPoolStart(t *testing.T) {
	task1 := NewTask(func() error {
		fmt.Println("execting task, it take second")
		time.Sleep(1 * time.Second)
		return nil
	})
	task2 := NewTask(func() error {
		fmt.Println("execting task, it take two second")
		time.Sleep(2 * time.Second)
		return nil
	})
	task3 := NewTask(func() error {
		fmt.Println("execting task, it take three second")
		time.Sleep(3 * time.Second)
		return nil
	})
	tasks := []*Task{
		task1,
		task2,
		task3,
	}

	p := NewPool(tasks, 2)
	p.Run()
}

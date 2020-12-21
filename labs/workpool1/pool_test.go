package workpool1

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	pool := NewPool(2)
	if pool.completed != false {
		t.Error("The Pool should be incomplete")
	}
}

func ResourceProcessor(resource interface{}) error {
	fmt.Printf("Resource processor got: %s", resource)
	fmt.Println()
	time.Sleep(time.Second)
	return nil
}

func ResultProcessor(result Result) error {
	fmt.Printf("Result processor got: %s", result.Err)
	fmt.Println()
	time.Sleep(time.Millisecond * 500)
	return nil
}

func TestPool_Start(t *testing.T) {
	strings := []string{"first", "second", "third", "fourth", "fifth", "sixth"}
	resources := make([]interface{}, len(strings))
	for i, s := range strings {
		resources[i] = s
	}

	pool := NewPool(2)
	pool.Start(resources, ResourceProcessor, ResultProcessor)
}

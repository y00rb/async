package compare

import (
	"sync"
	"time"
)

// 读多写少(读占 90%)
// 读少写多(读占 10%)
// 读写一致(各占 50%)

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

// const cost = time.Nanosecond * 100 // 时间降为 0.1 微秒
// const cost = time.Microsecond * 10 // 时间增加到 10 微秒

// 简单的互斥锁
type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

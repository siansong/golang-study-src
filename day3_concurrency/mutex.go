package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc ++
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value val
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

// MutexMain run
func MutexMain() {
	c := SafeCounter{v: make(map[string]int)}
	for i:=0; i < 1000; i++ {
		go c.Inc("hi")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("hi"))
}
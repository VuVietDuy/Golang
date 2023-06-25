package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Mutex() {
	counter := &Counter{}

	for i := 0; i < 100; i++ {
		go func() {
			counter.Increment()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(counter.GetValue())
}

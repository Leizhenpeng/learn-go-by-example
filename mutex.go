package main

import (
	"sync"
)

type container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *container) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[key]++
}

func main() {
	c := &container{
		counter: map[string]int{
			"foo": 1,
			"bar": 1,
		},
	}

	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			c.Inc(name)
		}
	}
	wg.Add(3)
	go doIncrement("foo", 10)
	go doIncrement("bar", 10)
	go doIncrement("bar", 10)
	wg.Wait()
	println(c.counter["foo"], c.counter["bar"])

}

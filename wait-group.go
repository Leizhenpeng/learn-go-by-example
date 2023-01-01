package main

import (
	"fmt"
	"sync"
	"time"
)

func newWorker(id int) {
	fmt.Println("worker", id, "started")
	time.Sleep(time.Second)
	fmt.Println("worker", id, "done")
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			newWorker(id)
		}(i)
	}
	wg.Wait()
}

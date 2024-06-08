package main

import (
	"fmt"
	"sync"
)

var (
	counter int // shared variable
	//mutex   sync.Mutex
)

// If we have a shared variable accessed by two goroutines, they could race to increment the counter. Therefore, we need a mutex to ensure that only one goroutine can add to or read the value at a time.

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			// mutex.Lock()
			counter++
			// mutex.Unlock()
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			// mutex.Lock()
			counter++
			// mutex.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

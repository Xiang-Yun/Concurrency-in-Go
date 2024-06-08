package main

import (
	"fmt"
	"sync"
)

var (
	// shared variable
	counter = 0
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			counter++
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			counter++
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

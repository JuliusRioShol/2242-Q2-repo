// Wait groups in the Go Programming language
// We can use a waitGroup to wait for multiple goroutines to finish.
package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Println("Go routine 1: hello")
	wg.Done()
}
func main() {
	// wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)

	wg.Wait() //we block until the counter is setback to 0;
	fmt.Println("Main has terminated")
}

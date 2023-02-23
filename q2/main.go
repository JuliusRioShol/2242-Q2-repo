// Go Routines:

package main

import (
	"fmt"
	"time"
)

// function: Prints numbers from 1-3 along  with the passed string.
func count(s string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, " : ", i)
	}
}
func main() {

	fmt.Println("Main started...")
	// Go routines
	go count("1st goroutine")
	go count("2nd goroutine")

	//wait for goroutine to finish before main goroutine ends:
	time.Sleep(time.Second)
	fmt.Println("...Main Finished")
}

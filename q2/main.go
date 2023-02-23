//Channels in Go :

package main

import "fmt"

func recieve(ch chan int) {
	x := <-ch

	fmt.Println(x)
}

func give(ch chan int) {
	ch <- 10

}

func finished(ch chan bool) {
	ch <- true
}

func main() {

	ch := make(chan int)
	go recieve(ch)
	ch <- 90 //nothing

	go give(ch)
	fmt.Println(<-ch)

	//x := <-ch
	fmt.Println("Main is back on track!")

	isDone := make(chan bool)
	go finished(isDone)

	if <-isDone {
		fmt.Println("Yeah ! function is done")
	}

	//Burffered channels
	uch := make(chan int, 3)

	uch <- 1
	uch <- 2
	uch <- 3

	for i := 0; i < 3; i++ {
		fmt.Println(<-uch)
	}

	uch <- 89
	fmt.Println(<-uch)
}

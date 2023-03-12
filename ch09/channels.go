package main

import "fmt"

func makeEventNums(count int, in chan<- int) {
	for i := 0; i < count; i++ {
		in <- 2 * i
	}
}

func main() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4
	close(ch)

	for i := 0; i < 4; i++ {
		if val, opened := <-ch; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}
}

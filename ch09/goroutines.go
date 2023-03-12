package main

import "fmt"

func count(start, stop, delta int) {
	for i := start; i < stop; i += delta {
		fmt.Println(i)
	}
}

func main() {
	go count(10, 50, 10)
	go count(60, 100, 10)
	go count(110, 200, 20)

	go func() {
		count(40, 60, 10)
	}()

	fmt.Scanln()
}

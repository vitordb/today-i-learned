package main

import "fmt"

func main() {
	ch := make(chan int)

	go publish(ch)

	for i := range ch {
		fmt.Printf(" Result: %d", i)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

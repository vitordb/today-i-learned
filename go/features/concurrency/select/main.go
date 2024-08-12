package main

import "time"

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		chan1 <- "Blue"
	}()

	select {
	case msg1 := <-chan1:
		println(msg1)
	case <-time.After(1 * time.Second):
		println("Timeout Blue")

	}

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Red"
	}()

	select {
	case msg2 := <-chan2:
		println(msg2)
	case <-time.After(3 * time.Second):
		println("Timeout Red")
	}

}

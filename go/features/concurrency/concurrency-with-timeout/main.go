package main

import (
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Ping"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			println(msg1)
		case msg2 := <-ch2:
			println(msg2)
		}
	}

}

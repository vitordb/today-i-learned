package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for x := range jobs {
		fmt.Printf("worker %d started job %d\n", id, x)
		time.Sleep(time.Second)
	}
}

func main() {

	data := make(chan int)
	qtdWorkers := 10

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data, nil)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func task(name string) {
	for range make([]struct{}, 10) {
		fmt.Printf("Task %s is running\n", name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	wg.Add(30)

	go task("A")

	go task("B")

	go func() {
		for range make([]struct{}, 10) {
			fmt.Printf("Task %s is running\n", "test")
			time.Sleep(1 * time.Second)
			wg.Done()
		}
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("id: %d, task name %s\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// thread 1
func main() {

	wg := sync.WaitGroup{}

	wg.Add(3)

	go task("C", &wg)
	go task("B", &wg)
	go task("A", &wg)

	wg.Wait()
}

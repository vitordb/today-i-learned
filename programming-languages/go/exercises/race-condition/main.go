package main

import (
	"fmt"
	"sync"
)

// var arrayTop []int

func main() {
	fmt.Println("comecando o playground")

	var (
		wg sync.WaitGroup
		n int
N = 10000
	)

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			n++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(n)
}

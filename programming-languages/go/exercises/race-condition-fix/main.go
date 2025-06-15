package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("comecando o playground")

	var (
		wg sync.WaitGroup
		n  int
		N  = 10000
		mu sync.Mutex
	)

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			mu.Lock()
			n++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(n)
}

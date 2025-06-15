package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello world"
	}()

	test := <-ch

	fmt.Println(test)
}

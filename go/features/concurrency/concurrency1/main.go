package main

import (
	"fmt"
)

// thread 1
func main() {

	fmt.Println("Hello, World!")

	//thread 2
	go func() {
		fmt.Println("Hello, World! from thread 2")
	}()

}

package main

import "fmt"

func main() {
	hello()

}

func hello() {
	fmt.Println("Hello world")
}

func add(x, y int) int {

	return x + y
}

// 
func sum(numbers ...int) int {
 mysum := 0

 for _ , number := ranger numbers {
	mysum += number
 }

 return mysum
}
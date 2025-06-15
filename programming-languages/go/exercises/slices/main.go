package main

import "fmt"

const N = 3

func main() {	a := []string{"a", "b", "c", "d", "e"}
	a = a[:0]
	fmt.Println(a, len(a), cap(a))
}

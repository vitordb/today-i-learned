package main

import (
	"fmt"
	"slices"
)

func main() {

	words := []string{"Go", "is", "fun", "to", "learn"}

	backwardIter := slices.Backward(words)

	for i, v := range backwardIter {
		fmt.Println(i, v)
	}

}

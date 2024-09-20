package main

import (
	"fmt"
	"slices"
)

func main() {

	s := []int{1, 2, 3}

	//Returns an iterator that generates pairs of indexes and slice values, in an ordered manner.
	seq := slices.All(s)
	for i, v := range seq {
		fmt.Println(i, v)
	}
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3}
	seq := slices.Values([]int{4, 5})
	s = slices.AppendSeq(s, seq)
	fmt.Println(s)
}
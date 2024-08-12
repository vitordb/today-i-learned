package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
	out:
		for j := 0; j < 3; j++ {
			fmt.Println(i, ",", j, " ")
			break out
		}

		fmt.Println("test")

	}

}

package main

import "fmt"

func main() {

	/*Basic for loop has three components separated by semcolons:
	  the init statement: executed before the first iteration
	  the condition expression: evaluated before every iteration
	  the post statement: eecuted at the end of every iteration*/

	for i := 0; i < 10; i++ {
		println("for loop", i)
	}

	/* Optional init and posts statements - same as while in other languages*/
	sum := 1

	for sum < 15 {
		sum += sum
		fmt.Println(sum)
	}

	/* for each range loop */
	strings := []string{"hello", "world"}
	for k, v := range strings {
		fmt.Println(k, v)
	}

	/* exit a loop */
	summ := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 {
			continue
		}

		sum += i

	}
	fmt.Println(summ)

}

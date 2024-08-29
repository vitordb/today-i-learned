package main

import "fmt"

func publisher(ch chan<- map[string]int) {
	for i := 0; i < 10; i++ {
		ch <- map[string]int{"key": i}
	}
	close(ch)

}

func reader(ch <-chan map[string]int) {
	mapExample := make(map[string]int)

	for k := range ch {
		for key, value := range k {
			mapExample[key] = value
			fmt.Println(mapExample)
		}
	}

}

func main() {

	ch := make(chan map[string]int)

	go publisher(ch)

	reader(ch)
}

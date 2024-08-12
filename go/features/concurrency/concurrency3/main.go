package main

import (
	"fmt"
	"net/http"
)

var numberOfRequests int

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		numberOfRequests++
		fmt.Println("Number of requests: ", numberOfRequests)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("sever is running")

}

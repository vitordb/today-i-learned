package main

import "net/http"

type home struct {
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

	mux.Handle("/", home{})

}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

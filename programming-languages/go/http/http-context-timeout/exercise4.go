package main

import (
	"context"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      timeOutMiddleware(mux, 5*time.Second),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func timeOutMiddleware(next http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		finished := make(chan struct{})
		go func() {
			next.ServeHTTP(w, r.WithContext(ctx))
			close(finished)
		}()

		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				http.Error(w, "Timeout", http.StatusRequestTimeout)
			case context.Canceled:
				http.Error(w, "Request canceled", http.StatusRequestTimeout)

			}
		case <-finished:
			return

		}

	})
}

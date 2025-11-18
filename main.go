package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	HOST = "127.0.0.1"
	PORT = 8081
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		Addr:         fmt.Sprintf("%s:%d", HOST, PORT),
		Handler:      mux,
	}
	fmt.Printf("[LOG] server is starting at %s:%d\n", HOST, PORT)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("[ERROR] server failed to start: %s\n", err)
	}
}

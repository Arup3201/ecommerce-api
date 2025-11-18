package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Arup3201/ec-api/models"
	"github.com/Arup3201/ec-api/routes"
)

const (
	HOST = "127.0.0.1"
	PORT = 8081
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/customers/login", routes.Customer.Login)

	models.InitDB("user:pass@localhost/bookstore")

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

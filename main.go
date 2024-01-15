package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a route handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is the response from the server!")
	}

	// Register the route handler function with the default ServeMux
	http.HandleFunc("/", handler)
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World !")
	})

	// Specify the port to listen on
	port := 3000

	// Start the server and listen for incoming requests
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

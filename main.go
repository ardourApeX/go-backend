package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// Specify the default port
	defaultPort := 3000

	// Parse the port flag
	portPtr := flag.Int("port", defaultPort, "server port to listen on")
	flag.Parse()

	// Start the server and listen for incoming requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Construct the server address
	serverAddr := fmt.Sprintf(":%d", *portPtr)

	fmt.Printf("Server is running on %s...\n", serverAddr)

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}

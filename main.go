package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Parse the port flag
	portPtr := flag.String("port", os.Getenv("PORT"), "server address to listen on")
	flag.Parse()

	// Start the server and listen for incoming requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Construct the server address
	fmt.Println("This is the port: ", *portPtr)
	serverAddr := fmt.Sprintf("0.0.0.0:%s", *portPtr)

	fmt.Printf("Server is running on %s...\n", serverAddr)

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}

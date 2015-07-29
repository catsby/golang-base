package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(response http.ResponseWriter, request *http.Request) {
	log.Printf("-- request received")
	fmt.Fprintf(response, "Hello –, %s!", request.URL.Path[1:])
}

func main() {

	log.Printf("Main server starting up")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

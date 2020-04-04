package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	text := os.Getenv("APP_TEXT")
	hostname, _ := os.Hostname()
	if text == "" {
		text = hostname
	}

	// Get request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Fprintf(w, "Incorrect method. GET only.\n")
			return
		}
		fmt.Fprintf(w, "Hello world\n")
		fmt.Fprintf(w, "text: %s\n", text)
		fmt.Fprintf(w, "hostname: %s\n", hostname)
	})

	var err error
	port = ":" + port
	fmt.Printf("server(http) %s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

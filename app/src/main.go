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
	if text == "" {
		hostname, _ := os.Hostname()
		text = hostname
	}

	// Get request
	// curl http://localhost:8080/get?query=aaa\&test=name\&name=xshoji
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// - [How can I handle http requests of different methods to / in Go? - Stack Overflow](https://stackoverflow.com/questions/15240884/how-can-i-handle-http-requests-of-different-methods-to-in-go)
		if r.Method != "GET" {
			fmt.Fprintf(w, "Incorrect method. GET only.\n")
			return
		}

		fmt.Fprintf(w, "Hello world\n")
		fmt.Fprintf(w, "text: %s\n", text)
	})

	var err error
	port = ":" + port
	fmt.Printf("server(http) %s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

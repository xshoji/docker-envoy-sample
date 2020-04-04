package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// get text
	text := os.Getenv("APP_TEXT")
	// get hostname
	hostname, _ := os.Hostname()
	// get port
	port := os.Getenv("APP_PORT")
	if port == "" {
		// default listen port
		port = "8080"
	}

	// GET request handling
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Incorrect http method. GET only.\n")
			return
		}
		// output an access log to /dev/stdout
		log.Print("| RemoteAddr: ", r.RemoteAddr, ", RequestURI:", r.RequestURI, ", UserAgent:", r.UserAgent())
		// response
		fmt.Fprintf(w, "Hello world\n")
		fmt.Fprintf(w, "text: %s\n", text)
		fmt.Fprintf(w, "hostname: %s\n", hostname)
	})

	// start server
	fmt.Printf("server(http) %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

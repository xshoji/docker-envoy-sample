package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// get proxyDestination ( http://hostname:port )
	proxyDestination := os.Getenv("FRONT_PROXY_DESTINATION")
	if proxyDestination == "" {
		log.Fatal("FRONT_PROXY_DESTINATION is not defined.")
	}
	// get port
	port := os.Getenv("FRONT_PORT")
	if port == "" {
		// default listen port
		port = "8080"
	}

	// GET request handling
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// output an access log to /dev/stdout
		log.Print("| RemoteAddr:", r.RemoteAddr, ", RequestURI:", r.RequestURI, ", UserAgent:", r.UserAgent())
		// response
		response, err := DoHttpGet(proxyDestination, map[string]string{})
		if err != nil {
			log.Fatal(err)
		}
		response.Write(w)
	})

	// start server
	fmt.Printf("server(http) %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//=======================================
// HTTP Utils
//=======================================

// Get request
func DoHttpGet(url string, headers map[string]string) (*http.Response, error) {
	// GET
	log.Println(">---------- Get request start ---------->")
	log.Printf("url : %v\n", url)
	resp, err := DoHttpRequest("GET", url, nil, headers)
	log.Println("<----------  Get request end  ----------<")
	return resp, err
}

func DoHttpRequest(method string, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for key, value := range headers {
		log.Printf("header [%s] : %s\n", key, value)
		req.Header.Set(key, value)
	}
	return http.DefaultClient.Do(req)
}

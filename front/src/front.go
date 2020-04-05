package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	origin, _ := url.Parse(proxyDestination)
	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// output an access log to /dev/stdout
		log.Print("| RemoteAddr:", r.RemoteAddr, ", RequestURI:", r.RequestURI, ", UserAgent:", r.UserAgent())
		proxy.ServeHTTP(w, r)
	})

	// start server
	fmt.Printf("server(http) %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

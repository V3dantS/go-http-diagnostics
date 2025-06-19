package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func requestLogger(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	fmt.Println("========== Request Info ==========")
	fmt.Printf("Time:           %s\n", start.Format(time.RFC3339))
	fmt.Printf("Protocol:       %s\n", r.Proto)
	fmt.Printf("Method:         %s\n", r.Method)
	fmt.Printf("URL:            %s\n", r.URL.String())
	fmt.Printf("Remote Address: %s\n", r.RemoteAddr)
	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}
	fmt.Println("=================================")

	// Set response headers
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Server-Protocol", r.Proto)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Go HTTP Server!\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestLogger)

	// TLS server (required for HTTP/2)
	server := &http.Server{
		Addr:         ":8443",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting diagnostic server on https://localhost:8443")
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}

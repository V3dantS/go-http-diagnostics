package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func makeRequest(url string, useHTTP2 bool, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, // allow self-signed certs
		ForceAttemptHTTP2: useHTTP2,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	// Create request with custom headers
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		fmt.Printf("[Client %d] Request error: %v\n", id, err)
		return
	}

	req.Header.Set("User-Agent", "Go-HTTP-Diagnostics")
	req.Header.Set("X-Debug-Client-ID", fmt.Sprintf("client-%d", id))

	// Send request and record time
	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("[Client %d] Request failed: %v\n", id, err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("[Client %d] Protocol: %s | Time: %v | Status: %d\n",
		id, resp.Proto, elapsed, resp.StatusCode)
	fmt.Printf("[Client %d] Response: %s\n", id, string(body))
}

func main() {
	url := "https://localhost:8443"

	fmt.Println("=== HTTP/1.1 Test ===")
	var wg1 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg1.Add(1)
		go makeRequest(url, false, i, &wg1)
	}
	wg1.Wait()

	time.Sleep(2 * time.Second)

	fmt.Println("\n=== HTTP/2 Test ===")
	var wg2 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go makeRequest(url, true, i, &wg2)
	}
	wg2.Wait()
}

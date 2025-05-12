package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type headerList map[string]string

func (h *headerList) String() string {
	return fmt.Sprintf("%v", *h)
}

func (h *headerList) Set(value string) error {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid header format, expected Key: Value")
	}
	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	(*h)[key] = val
	return nil
}

var (
	port      = flag.String("port", "8080", "Port to listen on")
	targetURL = flag.String("target", "", "Target domain to forward requests to (e.g., https://example.com)")
	headers   = headerList{}
)

func handler(w http.ResponseWriter, r *http.Request) {
	if *targetURL == "" {
		http.Error(w, "Target domain not configured", http.StatusInternalServerError)
		return
	}

	// Build target URL with query params
	target := *targetURL + r.URL.Path
	if r.URL.RawQuery != "" {
		target += "?" + r.URL.RawQuery
	}

	// Forward original request
	req, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Forward initial headers if present
	for k, vv := range r.Header {
		for _, v := range vv {
			req.Header.Add(k, v)
		}
	}

	// Add custom headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Upstream error: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy headers and status code
	for k, v := range resp.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	flag.Var(&headers, "H", "Custom header to add (repeatable, format: 'Key: Value')")
	flag.Parse()

	http.HandleFunc("/", handler)
	addr := ":" + *port
	fmt.Printf("Listening on port %s and forwarding to %s\n", *port, *targetURL)
	log.Fatal(http.ListenAndServe(addr, nil))
}

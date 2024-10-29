package proxy

import (
	"focus-guard/schedule"
	"log"
	"net/http"
	"net/http/httputil"
)

// StartProxy starts the HTTP/HTTPS proxy server
func StartProxy() error {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			if schedule.IsBlocked(req.URL.Hostname()) {
				log.Printf("Blocked URL: %s", req.URL)
				req.URL = nil // This will cause the proxy to return a 502 Bad Gateway
				return
			}

			// Set the scheme and host for the outgoing request
			req.URL.Scheme = "http"
			if req.TLS != nil {
				req.URL.Scheme = "https"
			}
			req.URL.Host = req.Host
		},
	}

	log.Println("Starting proxy server on :8080")
	return http.ListenAndServe(":8080", proxy)
}

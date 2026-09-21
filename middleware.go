package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		fmt.Printf("[Middleware] Received Request | Method: %s | Path: %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
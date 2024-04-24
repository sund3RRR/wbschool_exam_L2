package main

import (
	"fmt"
	"net/http"
)

// Middleware logger
func loggingWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logStr := fmt.Sprintf("Request Info:\n"+
			"\tMethod: %s\n"+
			"\tPath: %s\n"+
			"\tRemote Address: %s", r.Method, r.URL.Path, r.RemoteAddr)

		fmt.Println(logStr)

		next.ServeHTTP(w, r)
	})
}

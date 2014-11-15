package middle

import (
	"log"
	"net/http"
)

// Very simple Console Logger
func SimpleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

// Check for a specific method
func CheckMethod(next http.Handler, m string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == m {
			next.ServeHTTP(rw, req)
		} else {
			http.Redirect(rw, req, "/", http.StatusBadRequest)
		}
	})
}

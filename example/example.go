package main

import (
	"github.com/squiidz/fur"
	"log"
	"net/http"
)

func main() {
	// Create a NewServer
	s := fur.NewServer("localhost", ":8080", false)

	// Set some default middleware for all the handlers
	s.Stack(MiddleLog)
	// Add a new routes and add some middleware for this one only
	s.AddRoute("/nuts", DefaultHandler)
	s.AddRoute("/", DefaultHandler, MiddleRedirect)
	s.AddStatic("/public/", ".")
	// Start the server
	s.Start()
}

// Application Handler
func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	
	rw.Write([]byte()
}

// Middleware Logger
func MiddleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fur.Set(req, "api", "123456789")
		log.Printf("[%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

// Useless Middleware, just for example
func MiddleRedirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.RequestURI != "/nuts" {
			http.Redirect(rw, req, "http://google.com", http.StatusFound)
		} else {
			next.ServeHTTP(rw, req)
		}
	})
}

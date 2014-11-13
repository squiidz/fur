package main

import (
	"github.com/squiidz/fur"
	"log"
	"net/http"
)

func main() {
	// Create a NewServer
	s := fur.NewServer("localhost", ":8080", false)
	// Set some default middleware for all the handler
	s.Stack(MiddleLog)
	// Add a new route and add some middleware for this one only
	s.AddRoute("/home", DefaultHandler)

	// Run the server
	s.Start()
}

// Application Handler
func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Welcome to my website"))
}

// Middleware Logger
func MiddleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

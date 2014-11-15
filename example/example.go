package main

import (
	"log"
	"net/http"

	"github.com/squiidz/fur"
	"github.com/squiidz/fur/context"
)

func main() {
	// Create a NewServer
	s := fur.NewServer("localhost", ":8080", false)

	// Set some default middleware for all Route
	s.Stack(MiddleLog)

	// Add a new routes and add some middleware for this one only
	// You can force a HTTP method (.Get(), .Post(), .Put(), .Delete())
	s.AddRoute("/nuts", DefaultHandler).Get()
	s.AddRoute("/", DefaultHandler, MiddleRedirect)
	s.AddStatic("/public/", "my/assets/folder/")
	
	// Start the server
	s.Start()
}

// Application Handler
func DefaultHandler(rw http.ResponseWriter, req *http.Request) {
	// Short Retrive Context way
	value := context.FindContext(req).Get("MyKey")
	
	rw.Write([]byte(value.(string))
}

// Middleware Logger
func MiddleLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Short Context way
		context.NewContext(req).Set("MyKey", "MyValue")

		log.Printf("[%s] %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

// Useless Middleware, just for example
func MiddleRedirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.RequestURI != "/nuts" {
			http.Redirect(rw, req, "/nuts", http.StatusFound)
		} else {
			next.ServeHTTP(rw, req)
		}
	})
}

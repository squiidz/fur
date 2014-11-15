/********************************
*** Web server API for Go     ***
*** Code is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/squiidz        ***
*********************************/

package fur

import (
	"fmt"
	"log"
	"net/http"
)

var (
	origin Origin
)

// Simple Server structure for a web server.
type Server struct {
	Host   string
	Port   string
	Log    bool
	Mux    *http.ServeMux
	Routes []*Route
}

type MiddleWare func(http.Handler) http.Handler

type Origin []MiddleWare

// Create a NewServer instance with the given value.
// Host: "localhost"
// Port: ":8080"
// Log: true/false
// Options: functions to run on the server instance who's gonna be return.
func NewServer(host string, port string, log bool, options ...func(s *Server)) *Server {
	svr := Server{host, port, log, http.NewServeMux(), []*Route{}}
	if options != nil {
		for _, option := range options {
			option(&svr)
		}
	}
	return &svr
}

// Add Global Middleware to origin
func (s *Server) Stack(middles ...MiddleWare) {
	for _, middle := range middles {
		origin = append(origin, middle)
	}
}

// Start Listening on host and port of the Server.
// Log the request if the log was initiated as true in NewServer.
func (s *Server) Start() {
	fmt.Printf("[+] Server Running on %s ... \n", s.Port)
	if s.Routes != nil {
		for _, r := range s.Routes {
			s.Mux.Handle(r.Path, r)
		}
	}
	if s.Log {
		http.ListenAndServe(s.Host+s.Port, s.logger(s.Mux))
	}
	http.ListenAndServe(s.Host+s.Port, s.Mux)
}

// Add function with the right sigature to the Server Mux
// and chain the provided middlewares on it.
func (s *Server) AddRoute(path string, f func(rw http.ResponseWriter, req *http.Request), middles ...MiddleWare) *Route {
	var stack http.Handler
	var midStack = origin

	if middles != nil || midStack != nil {
		for _, mid := range middles {
			midStack = append(midStack, mid)
		}
		stack = midStack[0](http.HandlerFunc(f))
		stack = wrap(stack, midStack[1:])
	} else {
		stack = http.HandlerFunc(f)
	}

	r := Route{path, stack, ""}
	s.Routes = append(s.Routes, &r)
	return &r
}

// Temporary way for serving static files
func (s *Server) AddStatic(path string, dir string) {
	s.Mux.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

// Log request to the Server.
func (s *Server) logger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
		mux.ServeHTTP(rw, req)
	})
}

// Only Wrap the middleware on the provided http.Handler
func wrap(stack http.Handler, middles []MiddleWare) http.Handler {
	for i := len(middles) - 1; i >= 0; i-- {
		stack = middles[i](stack)
	}

	return stack
}

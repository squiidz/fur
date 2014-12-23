/********************************
*** Web server API for Go     ***
*** Code is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/squiidz        ***
*********************************/

package fur

import (
	"fmt"
	"net/http"
	//"reflect"

	"github.com/squiidz/bone"
	"github.com/squiidz/claw"
)

type Plex interface {
	Handle(string, http.Handler)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// Simple Server structure for a web server.
type Server struct {
	Host   string
	Port   string
	mux    Plex
	global *claw.Claw
	routes []*bone.Route
}

// Create a NewServer instance with the given value.
// Host: "localhost"
// Port: ":8080"
// mux: Any Type which implement Plex (http.NewServeMux(), bone.NewMux() etc..)
// Options: functions to run on the server instance who's gonna be return.
func NewServer(host string, port string, p Plex, options ...func(s *Server)) *Server {
	svr := Server{host, port, p, claw.New(), []*bone.Route{}}
	if options != nil {
		for _, option := range options {
			option(&svr)
		}
	}
	return &svr
}

// Create a NewServer instance with the default http.NewServeMux().
// Host: "localhost"
// Port: ":8080"
// Options: functions to run on the server instance who's gonna be return.
func NewServerMux(host string, port string, options ...func(s *Server)) *Server {
	svr := Server{host, port, bone.New(), claw.New(), []*bone.Route{}}
	if options != nil {
		for _, option := range options {
			option(&svr)
		}
	}
	return &svr
}

// Add Global Middleware to origin
func (s *Server) Stack(middles ...interface{}) {
	s.global.Wrap(middles)
}

// Start Listening on host and port of the Server.
// Log the request if the log was initiated as true in NewServer.
func (s *Server) Start() {
	fmt.Printf("[+] Server listening on: %s\n", s.Port[1:])
	if s.routes != nil {
		for _, r := range s.routes {
			s.mux.Handle(r.Path, r)
		}
	}
	http.ListenAndServe(s.Host+s.Port, s.mux)
}

// Add function with the right sigature to the Server Mux
// and chain the provided middlewares on it.
func (s *Server) AddRoute(path string, f func(rw http.ResponseWriter, req *http.Request), middles ...interface{}) *bone.Route {
	var stack http.Handler

	if middles != nil {
		for i, m := range middles {
			if i == 0 {
				stack = s.global.Use(f).Add(m)
			}
			stack = s.global.Merge(stack).Add(m)
		}
	} else {
		stack = s.global.Use(f)
	}

	r := bone.NewRoute(path, stack)
	s.routes = append(s.routes, r)
	return r
}

// Temporary way for serving static files
func (s *Server) AddStatic(path string, dir string) {
	fileHandler := http.StripPrefix(path, http.FileServer(http.Dir(dir)))
	s.mux.Handle(path, fileHandler)
}

// Only Wrap the middleware on the provided http.Handler
func wrap(stack http.Handler, middles []claw.MiddleWare) http.Handler {
	for i := len(middles) - 1; i >= 0; i-- {
		stack = middles[i](stack)
	}

	return stack
}

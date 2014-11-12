/********************************
*** Web server API for Go     ***
*** Code is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/squiidz        ***
*********************************/

package furserv

import (
	"fmt"
	"log"
	"net/http"
)

// Simple Server structure for a web server.
type Server struct {
	Host string
	Port string
	Log  bool
	Mux  *http.ServeMux
}

type MiddleWare func(http.Handler) http.Handler

// Create a NewServer instance with the given value.
// Host: "localhost"
// Port: ":8080"
// Log: true/false
// Options: functions to run on the server instance who's gonna be return.
func NewServer(host string, port string, log bool, options ...func(s *Server)) *Server {
	svr := Server{host, port, log, http.NewServeMux()}
	for _, option := range options {
		option(&svr)
	}
	return &svr
}

// Start Listening on host and port of the Server.
// Log the request if the log was initiated as true in NewServer.
func (s *Server) Start() {
	fmt.Printf("[+] Server Running on %s ... \n", s.Port)
	if s.Log {
		http.ListenAndServe(s.Host+s.Port, s.logger(s.Mux))
	}
	http.ListenAndServe(s.Host+s.Port, s.Mux)
}

// Add function with the right sigature to the Server Mux
// and chain the provided middlewares on it.
func (s *Server) AddRoute(pat string, f func(rw http.ResponseWriter, req *http.Request), middles ...MiddleWare) {
	if middles != nil {
		var stack http.Handler
		for i := len(middles) - 1; i >= 0; i-- {
			if i == len(middles)-1 {
				stack = middles[i](http.HandlerFunc(f))
			} else {
				stack = middles[i](stack)
			}
		}
		s.Mux.Handle(pat, stack)
	} else {
		s.Mux.Handle(pat, http.HandlerFunc(f))
	}

}

// Log request received by the Server.
func (s *Server) logger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
		mux.ServeHTTP(rw, req)
	})
}
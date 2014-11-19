/********************************
*** Web server API for Go     ***
*** Code is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/squiidz        ***
*********************************/

package middle

import (
	"log"
	"net/http"
	"runtime/debug"
)

type handler func(rw http.ResponseWriter, req *http.Request)

// transform Normal handler into middleware
func Mutate(h handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			h(rw, req)
			next.ServeHTTP(rw, req)
		})
	}
}

// Very simple Console Logger
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Printf("\x1b[42m[%s]\x1b[0m %s %s", req.Method, req.RequestURI, req.RemoteAddr)
		next.ServeHTTP(rw, req)
	})
}

// Recovery Middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				stack := debug.Stack()
				log.Printf("PANIC: %s\n%s", err, stack)

			}
		}()
		next.ServeHTTP(rw, req)
	})
}

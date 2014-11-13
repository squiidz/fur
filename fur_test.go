package fur

import (
	"log"
	"net/http"
	"testing"
)

func TestAddRoute(t *testing.T) {
	s := NewServer("localhost", ":8080", true)
	s.AddRoute("/test", testHandler, testMiddle)
}

func TestStack(t *testing.T) {
	s := NewServer("localhost", ":8080", true)
	s.Stack(testMiddle)
	s.AddRoute("/test", testHandler)
	s.Start()
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("test"))
}

func testMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println("TEST MIDDLEWARE")
		next.ServeHTTP(rw, req)
	})
}

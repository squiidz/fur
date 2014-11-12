package furserv

import (
	"net/http"
	"testing"
)

var (
	data int
)

func TestAddHandler(t *testing.T) {
	s := NewServer("localhost", ":8080", true)
	s.AddHandler("/test", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("TEST"))
	},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				data = 1
				t.Log("Middle test")
				next.ServeHTTP(rw, req)
			})
		})

	s.Start()
}

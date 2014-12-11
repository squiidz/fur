package fur

import (
	"net/http"
	"strings"
)

type Mux struct {
	Routes   []*Route
	NotFound http.HandlerFunc
}

func NewMux() *Mux {
	return &Mux{}
}

func (m *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	reqPath := req.URL.Path
	reqPart := strings.Split(req.RequestURI, "/")
	reqSize := len(reqPart)

	if !valid(reqPath) {
		http.Redirect(rw, req, reqPath[:len(reqPath)-1], http.StatusMovedPermanently)
		return
	}

	for _, r := range m.Routes {

		if reqSize == r.Size {

			if r.pattern.Exist {
				if v, ok := r.matcher(req.RequestURI); ok {
					req.URL.RawQuery = v.Encode() + "&" + req.URL.RawQuery
					r.handler.ServeHTTP(rw, req)
					return
				}
			} else {
				if req.RequestURI == r.Path {
					r.handler.ServeHTTP(rw, req)
					return
				}
			}
		}
	}

	if m.NotFound != nil {
		m.NotFound(rw, req)
	} else {
		http.NotFound(rw, req)
	}
}

func valid(path string) bool {
	pathLen := len(path)

	if pathLen > 1 && path[pathLen-1:] == "/" {
		return false
	}
	return true
}

func (m *Mux) Handle(s string, h http.Handler) {
	m.Routes = append(m.Routes, NewRoute(s, h))
}

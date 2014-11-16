/********************************
*** Web server API for Go     ***
*** Code is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/squiidz        ***
*********************************/

package fur

import (
	"net/http"
)

type Route struct {
	Path    string
	handler http.Handler
	Method  string
}

func NewRoute(url string, h http.Handler, m string) *Route {
	return &Route{url, h, m}
}

func (r *Route) Get() {
	r.Method = "GET"
}

func (r *Route) Post() {
	r.Method = "POST"
}

func (r *Route) Put() {
	r.Method = "PUT"
}

func (r *Route) Delete() {
	r.Method = "DELETE"
}

func (r Route) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if r.Method != "" {
		if req.Method == r.Method {
			r.handler.ServeHTTP(rw, req)
		} else {
			rw.WriteHeader(http.StatusBadRequest)
		}
	} else {
		r.handler.ServeHTTP(rw, req)
	}

}

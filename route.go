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

/*
var (
	action = []string{"POST", "GET", "PUT", "DELETE"}
)

func (r *Route) SetMethod(methods ...string) {
	var valid = []string{}
	if methods != nil {
		for _, a := range action {
			for _, m := range methods {
				if m == a {
					valid = append(valid, m)
				}
			}
		}
	}
}
*/

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
		//correct := methodValid(req.Method, r.Method)
		if req.Method == r.Method {
			r.handler.ServeHTTP(rw, req)
		} else {
			rw.WriteHeader(http.StatusBadRequest)
		}
	} else {
		r.handler.ServeHTTP(rw, req)
	}
}

/*
func methodValid(m string, valMe []string) bool {
	for _, v := range valMe {
		if m == v {
			return true
		}
	}
}
*/

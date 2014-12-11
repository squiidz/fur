/********************************
*** Web server API for Go     ***
*** Code is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/squiidz        ***
*********************************/

package context

import (
	"net/http"
	"sync"
)

// Context struct, req = the *http.Request provided with the NewContext func
type Context struct {
	req  *http.Request
	vars map[string]string
}

var (
	mutex sync.RWMutex
	memo  = make(map[*http.Request]Context)
)

// Create a new context struct, with vars map
func NewContext(req *http.Request) *Context {
	memo[req] = Context{req, make(map[string]string)}
	c := memo[req]
	return &c
}

// Retrive a context from the memo map and return.
func Find(req *http.Request) *Context {
	c := memo[req]
	if c.req != nil {
		con := memo[req]
		return &con
	}
	return nil
}

// Set a new key value pair in the context.Vars
func (c *Context) Set(key, value string) {
	mutex.Lock()
	c.vars[key] = value
	mutex.Unlock()
}

// Get a key value from a context
func (c *Context) Get(key string) string {
	return c.vars[key]
}

// Get the map of data of the context
func (c *Context) GetAll() map[string]string {
	return c.vars
}

// Delete a key
func (c *Context) Delete(key string) {
	delete(c.vars, key)
}

// Remove a Context from the memo map
func (c *Context) Destroy() {
	delete(memo, c.req)
}

package fur

import (
	"net/http"
	"sync"
)

// Context struct, req = the *http.Request provided with the NewContext func
type Context struct {
	req  *http.Request
	vars map[interface{}]interface{}
}

var (
	mutex sync.RWMutex
	memo  = make(map[*http.Request]Context)
)

// Create a new context struct, with vars map
func NewContext(req *http.Request) *Context {
	memo[req] = Context{req, make(map[interface{}]interface{})}
	c := memo[req]
	return &c
}

// Retrive a context from the memo map and return.
func FindContext(req *http.Request) *Context {
	con := memo[req]
	return &con
}

// Set a new key value pair in the context.Vars
func (c *Context) Set(key, value interface{}) {
	mutex.Lock()
	c.vars[key] = value
	mutex.Unlock()
}

// Get a key value from a context
func (c *Context) Get(key interface{}) interface{} {
	return c.vars[key]
}

// Get the map of data of the context
func (c *Context) GetAll() map[interface{}]interface{} {
	return c.vars
}

// Delete a key
func (c *Context) Delete(key interface{}) {
	delete(c.vars, key)
}

// Remove a Context from the memo map
func (c *Context) Destroy() {
	delete(memo, c.req)
}

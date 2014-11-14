package fur

import (
	"net/http"
	"sync"
)

var (
	mut     sync.RWMutex
	context = make(map[*http.Request]map[interface{}]interface{})
)

func Set(req *http.Request, key, value interface{}) {
	mut.Lock()
	if context[req] != nil {
		context[req][key] = value
	} else {
		context[req] = make(map[interface{}]interface{})
		context[req][key] = value
	}
	mut.Unlock()
}

func Get(req *http.Request, key interface{}) interface{} {
	mut.RLock()
	if context[req] != nil {
		return context[req][key]
	}
	mut.RUnlock()
	return nil
}

func Delete(req *http.Request, key interface{}) {
	mut.Lock()
	if context[req] != nil {
		delete(context[req], key)
	}
	mut.Unlock()
}

func Destroy(req *http.Request) {
	delete(context, req)
}

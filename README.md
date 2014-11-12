furServ
=======

## Simple way Go web server with Middleware chaining.

### Example
``` 
	package main
	
	import "github.com/squiidz/furserv"
	
	func main() {
	    server := furserv.NewServer("localhost", ":8080", true, option1, option2)

	    server.AddRoute("/home", HomeHandler, Middleware1, Middleware2, Middleware3)
	    server.AddRoute("/", DefaultHandler)

	    server.Start()
	}
```

### Middleware and Options
Every function who have ``` func (next http.Handler) http.Handler ``` can be pass as a MiddleWare.
Option need ``` func (s *Server) *Server ``` can be pass as a Options.

### Next Feature
- Easy static files serving instead of 
``` http.Handle("/root/", http.stripPrefix("/root/", http.FileServer(http.Dir("folder")))) ```
- Add General Middleware function too MiddleWare type. (The Middleware become activated on every handler by default)

### License
MIT

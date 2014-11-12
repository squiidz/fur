furServ
=======

## Simple way Go web server with middleware chaining.

### Example
``` 
	package main
	
	import "github.com/squiidz/furserv"
	
	server := furserv.NewServer("localhost", ":8080", true, option1, option2)

	server.AddRoute("/home", HomeHandler, Middleware1, Middleware2, Middleware3)
	server.AddRoute("/", DefaultHandler)

	server.Start()

```
### Next Feature
- easy static files serving 
instead of ``` http.Handle("/root/", http.stripPrefix("/root/", http.FileServer(http.Dir("folder")))) ```

### License
MIT

furServ
=======

### Simple way of making Go web server with middleware chaining.
``` 
	server := furserv.NewServer("localhost", ":8080", true)

	server.AddRoute("/home", HomeHandler, Middleware1, MiddleWare2, Middleware3)
	server.AddRoute("/", DefaultHandler)

	server.Start()

```
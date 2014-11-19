fur
=======

## What is fur ?

fur is just a simple web server api, to make things less repetitive.
No revolution of any kind in this package, just a simple way to build your
Web application basics. fur isn't trying to be a framework at all!
It is more of a toolkit (e.g [Gorilla](https://github.com/gorilla/mux)).

## Changes

- Remove the log argument from ` fur.NewServer() `, you can use your own logger or if you need a simple one you can use ` middle.Logger() `. You just have to use it in ` fur.Stack() `, if you want to use it on every handler.

- Now you can provide your own Multiplexer, you juste need to implement the Plex interface.
	If you want to use the default` http.ServeMux `, you can use ` fur.NewServerMux() `instead of ` fur.NewServer() `.

## Features

- Middleware Chaining.
- Global Middleware declaration.
- Shorter version, for static file serving.
- Server instance as a Struct.
- Force a HTTP method on a route.
- Chaining Options on the Server at Creation.
- Simple Context Struct

![alt tag](http://upload.wikimedia.org/wikipedia/commons/8/8c/Marmota.jpg)

## Example
```go
package main

import "github.com/squiidz/fur"

func main() {
	// You can use ` fur.NewServerMux() ` if you want the default http.ServeMux.
	server := fur.NewServer("localhost", ":8080", yourMux, option1, option2)
	// Set Global Middleware
	server.Stack(GlobalMiddleWare)

	// Serve Static files
	server.AddStatic("/public/", "../public")

	// Add multiple routes with chaining middleware and http Method
	server.AddRoute("/home", HomeHandler, Middleware1, Middleware2)
	server.AddRoute("/", DefaultHandler, MiddleWare3).Get()
	server.AddRoute("/data", DataHandler).Post()

	// Start Listening
	server.Start()
}
```

## Middlewares and Options
Every function that has the signature ` func (next http.Handler) http.Handler ` can be passed as a MiddleWare.
Or pass a ` func (rw http.ReponseWriter, req *http.Request) ` to ` middle.Mutate() `.

Every function that has the signature ` func (s *fur.Server) ` can be passed as a Option.

## Context

- Import ``` github.com/squiidz/fur/context ```
- Create a New Context with ``` cont := context.NewContext(req) ```
- Set some Key/Value ``` cont.Set("key", "value") ```
- Find the already created context ``` cont := context.FindContext(req) ```
- Retrive the key ``` cont.Get("key") ```

- Check the example folder if you want to see it in action.

## Next features
- Context Variables [75%]
- Shortway static files serving [DONE]
- Add Global Middleware [DONE]
- Force HTTP method on Route [DONE]
- Refactoring

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License
MIT

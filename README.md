fur
=======

## What is fur ?

fur is just a simple web server api, to make things less repetetive.
No revolution of any kind in this package, just a simple way to build your 
Web application basics. fur have not the goal to be a framework at all !
It is more in a way of toolkit (e.g Gorilla). 

## What it does ..?

- Middleware Chaining.
- Global Middleware declaration.
- Shorter version, for static files serving.
- Server instance as a Struct.
- Chaining Options on the Server at Creation.
- Simple Context Struct

![alt tag](http://upload.wikimedia.org/wikipedia/commons/8/8c/Marmota.jpg)

## Example
``` 
	package main
	
	import "github.com/squiidz/fur"
	
	func main() {
	    server := fur.NewServer("localhost", ":8080", true, option1, option2)

	    server.Stack(GlobalMiddleWare)

	    server.AddStatic("/public/", "../public")
	    server.AddRoute("/home", HomeHandler, Middleware1, Middleware2)
	    server.AddRoute("/", DefaultHandler, MiddleWare3)

	    server.Start()
	}
```

## Middlewares and Options
- Every function who have ` func (next http.Handler) http.Handler ` can be pass as a MiddleWare.

- Option siganture is ` func (s *fur.Server) ` every function with this one, can be pass as a Option.

## Context

- Create a New Context with ``` cont := fur.NewContext(req) ```
- Set some Key/Value ``` cont.Set("key", "value") ```
- Find the already created context ``` cont := fur.FindContext(req) ```
- Retrive the key ``` cont.Get("key") ``` 

- Check the example if you need a using case.

## Next Feature
- Context Variables [50%]
- Shortway static files serving [DONE] 
- Add Global Middleware [DONE]

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License
MIT

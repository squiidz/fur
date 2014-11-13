fur
=======

## Example
``` 
	package main
	
	import "github.com/squiidz/fur"
	
	func main() {
	    server := fur.NewServer("localhost", ":8080", true, option1, option2)

	    server.Stack(GlobalMiddleWare)

	    server.AddStatic("/public/", ".")
	    server.AddRoute("/home", HomeHandler, Middleware1, Middleware2)
	    server.AddRoute("/", DefaultHandler, MiddleWare3)

	    server.Start()
	}
```

## Middlewares and Options
- Every function who have ` func (next http.Handler) http.Handler ` can be pass as a MiddleWare.

- Option siganture is ` func (s *fur.Server) *fur.Server ` every function with this one, can be pass as a Option.

## Next Feature
- Context Variables
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

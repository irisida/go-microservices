![](/assets/microservicesWithGo.png)


# Basic WebServer with Go

Below we have an example of a trivial webserver with `Go`. Although this is not something that we would use in production it's useful for demonstration purposes in the sense that it covers the very basics of implementing a web server in `Go` using just the `net/http` package and the `HandleFunc` method.

There are many available frameworks written for implementing webservices in Go and the commonality is that underneath them all you'll find basics with the `net/http` package such as what we can see here. Essentially abstracting and wrapping the core functionality with utility functions and method naming conventions.

![](/01-introduction/assets/basic-server.png)

```shell
> go run main.go
```
Once we're running you can go to `localhost:8080/confirm` and you will see the web-server in action, albeit with a rather limited functionality.

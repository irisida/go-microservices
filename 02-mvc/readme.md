![](/assets/microservicesWithGo.png)

# The MVC pattern

We're going to layout/mock an MVC pattern project in this section and to get started we can setup the directories at our project root for how we will organise our source files.

![](/02-mvc/assets/mvc-layout-dirs.png)

## Getting the first endpoint active

To get our application enabled we need the entry point for the app, the basic server and a `HandleFunc` that will call the mapped function once that endpoint is hit on a browser or `curl` request.

So, let's see how that looks in our files at this stage.

![](/02-mvc/assets/mvc-main-01.png)
![](/02-mvc/assets/mvc-app-01.png)
![](/02-mvc/assets/mvc-users-controller-01.png)

```shell
> go run main.go
```
In another terminal we can run the following:
```shell
> curl localhost:8080/users -v
```
and we can expert a response along the lines of What we see here.
```shell
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /users HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Thu, 26 Nov 2020 22:27:02 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
* Closing connection 0
```

## Implementing the MVC on our endpoint

Now we've seen we can get the webserver running and accept a query param being passed in we need to implement our endpoint. To do so we'll create a mock database with a map of faked users in the code for now. The purpose here is to see the layers in action an the endpoint returning data that matches query params as well as handling errors.

Let's have a look at what we'll be implementing here as a concept.

![](/02-mvc/assets/mvc-flow-model.png)
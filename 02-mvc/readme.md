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

## But why the app.go file?

Fair question, why are we calling a function from the main.go file that has the actual main application logic in it in the `app.go` file when we could simply move that logic to the `main.go` file and call it from there? In short, because it makes for easier testing. If we have everything in the main we have to call the main in the testing which effectively starts our application. When this logic is residing in another file we can call that logic using mocks etc.

Might seem trivial at face value but it's a very useful pattern to have the `Go` predefined application entry point simply call a use defined application entry point.

## Implementing the MVC on our endpoint

Now we've seen we can get the webserver running and accept a query param being passed in we need to implement our endpoint. To do so we'll create a mock database with a map of faked users in the code for now. The purpose here is to see the layers in action an the endpoint returning data that matches query params as well as handling errors.

Let's have a look at what we'll be implementing here as a concept.

![](/02-mvc/assets/mvc-flow-model.png)

So let's see how that looks in terms of code changes. We're adding in the service wrapper and the `dao` (data access object) to our codebase to keep with the pattern and make sure we have clear separation of concerns well defined. Within the model layer itself you will see we will have the models themselves and the data access objects that use the models. While coupled, these form a single layer, here called the domain, but is the `M` in the `MVC` for model.

The `main.go` and the `app.go` remain unchanged, so I will omit further screenshots here to save having to digest them again. Let's see how the controller has changed. The controller implementation is updated to call the service for `GetUser` and handle the returned user data and pss it as the response to our `get` request to the endpoint. Where the request was invalid or bad it should also handle the error as the response too.

![](/02-mvc/assets/mvc-controller-02.png)

We can see the service is pretty transparent in that it is simply an abstract to the functionality of the `dao` in the domain layer. There's no much to add as its a straight through-and-through.

![](/02-mvc/assets/mvc-service-02.png)

We can then see in the `DAO` that we are encapsulating the data layer concerns and this would allow us to make the persistence layer interchangeable if required.

![](/02-mvc/assets/mvc-dao-02.png)

When running we can see the following for a condition for which we have no data.
```shell
❯ curl "localhost:8080/users?id=1234"
user 1234 was not found%
```

For conditions that do exist, ie, user IDs `123` and `456`.
```shell
❯ curl "localhost:8080/users?id=456"
{"ID":456,"Fname":"Four","Lname":"Fivesix","Email":"ol456@ohaye.oi"}%

❯ curl "localhost:8080/users?id=123"
{"ID":123,"Fname":"One","Lname":"Twothree","Email":"big123@wee123.net"}%
```

and for handling the badly formatted queries in the URL
```shell
❯ curl "localhost:8080/users?id=45ED"
user id is not the correct format%
```

## Adding an Error handler

Next improvement is to add an `ApplicationError` struct and we can handle errors in a more uniform way.

![](/02-mvc/assets/mvc-app-err.png)

Then we can update the sources using the error to take a pointer to this standardised type. We can see the demoed in the updated controller code below.

![](/02-mvc/assets/mvc-err-updated.png)

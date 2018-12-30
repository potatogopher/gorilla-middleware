# Chaining Middleware with Gorilla Mux
This little API shows you how you can use Gorilla's `Router.Use` function to register middlewares. This approach is an alternative to wrapping middlewares around middlewares and so on.

## Setup
```zsh
$ git clone git@github.com:potatogopher/gorilla-middleware.git
$ cd gorilla-middleware
$ go run main.go
```

## Middlewares

**Logger**

The logging middleware will log data about the requests being made by a client.

**CORS**

CORS provides Cross-Origin Resource Sharing. CORS will be handled for all routes that recognize the `OPTIONS` method. [More info](https://github.com/gorilla/mux/issues/381)

**Recovery Handler**

The recovery handler will prevent a panic from happening. It will end up responding with a 500 Internal Server Error.

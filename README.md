# go-router

## Introduction

This Git documentation provides an overview of the "httpRouting" package, which is designed to facilitate HTTP routing in Go applications. The package contains three main files: "router.go," "route.go," and "middleware.go."

## Files Overview

### router.go

- Defines the `router` struct, which manages the routing logic and CORS headers.
- Provides methods to configure CORS headers: `SetAllowOrigin`, `SetAllowMethods`, `SetAllowHeaders`, and `SetCredantials`.
- Implements route creation with `NewRoute` and request handling with `Serve` and `ServeWithCORS` functions.
- Provides utility functions to retrieve URL parameters from requests.

### route.go

- Defines the `route` struct, which represents a single route.
- Contains fields for the regular expression, route handler, and middleware functions.

### middleware.go

- Defines the `Middleware` function type used for HTTP middleware.
- Contains an example of a middleware function called `HasAccess`.

## Usage Example

The provided "main.go" file demonstrates the usage of the "httpRouting" package to create a simple HTTP server with a single route and a middleware function.

```go
func main() {
	r := httpRouting.NewRouterBuilder().
		SetAllowOrigin("http://localhost:3000").
		SetAllowMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}).
		SetAllowHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token"}).
		SetCredantials(true)

	r.NewRoute("get", `/home/(?P<id>\d+)`, <Handler>, ...<Middlewares>)

	http.HandleFunc("/", r.ServeWithCORS())

	log.Fatal(http.ListenAndServe(":8080", nil))
}


```

## Installation

To use the "httpRouting" package in your Go project, follow these steps:

1. Ensure you have Go installed on your machine.
2. Run the following command to add the package to your project:

```bash
go get github.com/Zewasik/go-router@v0.2.0
```

## Getting Started

After installing the package, you can use it in your Go code by importing it as follows:

```go
import httpRouting "github.com/Zewasik/go-router"
```

## Router Configuration

To create a new router instance and configure CORS headers, use the `NewRouterBuilder` function along with the provided methods:

```go
r := httpRouting.NewRouterBuilder().
    SetAllowOrigin("http://localhost:3000").
    SetAllowMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}).
    SetAllowHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token"}).
    SetCredantials(true)
```

## Route Creation

To add a new route to the router, use the `NewRoute` method. It requires the HTTP method, a regular expression, the handler function, and optional middleware functions:

```go
r.NewRoute("get", `/home/(?P<id>\d+)`, <HandleFunc>, ...<Middleware>)
```

## Middleware

Middleware functions can be used to perform pre-processing or post-processing tasks for requests. The example includes a middleware function called `HasAccess`:

```go
func MiddlewareExample() httpRouting.Middleware {
    return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// your code here

			next.ServeHTTP(w, r)
		})
	}
}
```

## Running the Server

To start the HTTP server, use the `http.ListenAndServe` function:

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

## Conclusion

The "httpRouting" package simplifies HTTP routing in Go applications by providing a flexible and easy-to-use router. Its modular design allows for easy integration of middleware functions to handle various request processing tasks.

This documentation provides an overview of the package's main components and usage examples. For more detailed information, you can refer to the comments and code in the provided files.

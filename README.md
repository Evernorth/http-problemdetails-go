# http-problemdetails-go

[![Go Reference](https://pkg.go.dev/badge/github.com/Evernorth/http-problemdetails-go.svg)](https://pkg.go.dev/github.com/Evernorth/http-problemdetails-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Evernorth/http-problemdetails-go)](https://goreportcard.com/report/github.com/Evernorth/http-problemdetails-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/Evernorth/http-problemdetails-go)](https://gtihub.com/Evernorth/http-problemdetails-go/releases)

## Description
Package `problemdetails` is an implementation of [IETF RFC 9457 Problem Details for HTTP APIs](https://www.rfc-editor.org/rfc/rfc9457.html), which is a specification for a standard error structure for HTTP APIs.

## Installation
```go get -u github.com/Evernorth/http-problemdetails-go```

## Features
 - Supports the following HTTP status codes:

| HTTP Code                 | Method                                                         |
|---------------------------|----------------------------------------------------------------|
| 400 Bad Request           | [`NewBadRequest()`](problemdetails/problemdetails.go)          |
| 401 Unauthorized          | [`NewUnauthorized()`](problemdetails/problemdetails.go)        |
| 403 Forbidden             | [`NewForbidden()`](problemdetails/problemdetails.go)           |
| 404 Not Found             | [`NewNotFound()`](problemdetails/problemdetails.go)            |
| 405 Method Not Allowed    | [`NewMethodNotAllowed()`](problemdetails/problemdetails.go)    |
| 409 Conflict              | [`NewConflict()`](problemdetails/problemdetails.go)            |
| 429 Too Many Requests     | [`NewTooManyRequests()`](problemdetails/problemdetails.go)     |
| 500 Internal Server Error | [`NewInternalServerError()`](problemdetails/problemdetails.go) |
| 502 Bad Gateway           | [`NewBadGateway()`](problemdetails/problemdetails.go)          |
| 503 Service Unavailable   | [`NewServiceUnavailable()`](problemdetails/problemdetails.go)  |
| 504 Gateway Timeout       | [`NewGatewayTimeout()`](problemdetails/problemdetails.go)      |


>Note: MIME type must be set in the response header to comply with the RFC 9457 specification.  The `MIMETypeJSON` constant is provided for this purpose.

## Usage
### Creating a simple HTTP ProblemDetails
The example code below demonstrates how to create simple Problem instances as well as how to add information to them using the With* functions.

You can create a Problem instance by simply calling functions like `problemdetails.NewInternalServerError()`, `problemdetails.NewNotFound()` and `problemdetails.NewBadRequest()`.

You can easily override the default Problem title and detail fields and provide custom extension fields by using the `WithTitle`, `WithDetail` and `WithExtension` functions.
```
package main

import (
	"fmt"
	"github.com/Evernorth/http-problemdetails-go/problemdetails"
	"github.com/go-chi/render"
	"net/http"
)

func handler(httpRespWriter http.ResponseWriter, httpReq *http.Request) {

	// If the request is not a GET, return a 405 Method Not Allowed
	if httpReq.Method != http.MethodGet {
		httpRespWriter.Header().Set("Content-Type", problemdetails.MIMETypeJSON)
		httpRespWriter.WriteHeader(http.StatusMethodNotAllowed)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewMethodNotAllowed())
		return
	}

	// Get the example-type from the query parameters
	queryParams := httpReq.URL.Query()
	exampleType := queryParams.Get("example-type")

	if exampleType == "basic" {

		// Return a 500 Internal Server Error, using the NewInternalServerError function.
		httpRespWriter.Header().Set("Content-Type", problemdetails.MIMETypeJSON)
		httpRespWriter.WriteHeader(http.StatusInternalServerError)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError())

	} else if exampleType == "advanced" {

		// Return a 500 Internal Server Error, using the NewInternalServerError and With* functions.
		httpRespWriter.Header().Set("Content-Type", problemdetails.MIMETypeJSON)
		httpRespWriter.WriteHeader(http.StatusInternalServerError)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError().
			WithTitle("KABOOM!!!").
			WithDetail("The unthinkable has occurred.").
			WithExtension("example1", "test").
			WithExtension("example2", "this could be a struct if you like"))

	} else {

		// Return a 400 Bad Request, using the NewBadRequest function.
		httpRespWriter.Header().Set("Content-Type", problemdetails.MIMETypeJSON)
		httpRespWriter.WriteHeader(http.StatusBadRequest)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewBadRequest().
			WithDetail("example-type must be basic or advanced."))

	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
```

Here is an example of the output from the above code:
```
HTTP/1.1 400 Bad Request
Content-Type: application/problem+json
Date: Tue, 15 Oct 2024 16:01:11 GMT
Content-Length: 263

{
  "type": "urn:problems:bad-request",
  "status": 400,
  "title": "Request could not be processed because it is invalid.",
  "detail": "example-type must be basic or advanced.",
  "instance": "urn:uuid:d9c65ded-ec53-464c-8b3a-be3e8c25bf67",
  "created": "2024-10-15T16:01:11.601188Z"
}
```

## Dependencies
See the [go.mod](go.mod) file.

## Support
If you have questions, concerns, bug reports, etc. See [CONTRIBUTING](CONTRIBUTING.md).

## License
http-problemdetails-go is Open Source software released under the [Apache 2.0 license](https://www.apache.org/licenses/LICENSE-2.0.html).

## Original Contributors
- Steve Sefton, Evernorth
- Shellee Stewart, Evernorth
- Cindy Chen, Evernorth

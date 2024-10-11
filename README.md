# http-problemdetails-go

An implementation of [IETF RFC 9457 Problem Details for HTTP APIs](https://www.rfc-editor.org/rfc/rfc9457.html) which is a specification for a standard error structure for HTTP APIs.


## Dependencies
See the [go.mod](go.mod) file.

## Building from Source

Detailed instructions on how to build the project from source. Also note where to get pre-built distribution, if building is not required.

## Installation
```go get -u github.com/Evernorth/http-problemdetails-go```

## Features
 - The MIME type of the response is `application/problem+json`.
 - Supports the following HTTP status codes:

| HTTP Code                           | Method                                          |
|-------------------------------------|-------------------------------------------------|
| 400 Bad Request                     | [`NewBadRequest()`](problemdetails.go)          |
| 401 Unauthorized                    | [`NewUnauthorized()`](problemdetails.go)        |
| 403 Forbidden                       | [`NewForbidden()`](problemdetails.go)           |
| 404 Not Found                       | [`NewNotFound()`](problemdetails.go)            |
| 409 Conflict                        | [`NewConflict()`](problemdetails.go)            |
| 429 Too Many Requests               | [`NewTooManyRequests()`](problemdetails.go)     |
| 500 Internal Server Error           | [`NewInternalServerError()`](problemdetails.go) |
| 501 Not Implemented                 | [`NewNotImplemented()`](problemdetails.go)      |
| 502 Bad Gateway                     | [`NewBadGateway()`](problemdetails.go)          |
| 503 Service Unavailable             | [`NewServiceUnavailable()`](problemdetails.go)  |
| 504 Gateway Timeout                 | [`NewGatewayTimeout()`](problemdetails.go)      |


>Note: MIME type must be set in the response header to comply with the RFC 9457 specification.

## Usage
### Creating a simple HTTP ProblemDetails
You can create a ProblemDetails instance by simply calling functions like `problemdetails.NewInternalServerError()`, `problemdetails.NewNotFound()` and `problemdetails.NewBadRequest()`.
```
package main

import (
    "github.com/Evernorth/http-problemdetails-go/problemdetails"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    "net/http"
)

func getSomething(httpRespWriter http.ResponseWriter, httpReq *http.Request) {
    
    err := doSomething()
    if err != nil {
        httpRespWriter.WriteHeader(http.StatusInternalServerError)
        render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError())
        httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
      return
    }
    ...
}
```
### Creating a complex HTTP ProblemDetails
You can easily override the default ProblemDetails title and detail fields and provide custom extension fields by using the `WithTitle`, `WithDetail` and `WithExtension` functions.
```
package main

import (
    "github.com/Evernorth/http-problemdetails-go/problemdetails"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    "net/http"
)

func getSomething(httpRespWriter http.ResponseWriter, httpReq *http.Request) {
    
    err := doSomething()
    if err != nil {
        httpRespWriter.WriteHeader(http.StatusInternalServerError)
        render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError().WithTitle(
            "KABOOM!!!").WithDetail("The unthinkable has occurred.").WithExtension(
                "example1", "test").WithExtension(
                    "example2", "this could be a struct if you like"))
        httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
        
        return
    }
    ...
}
```

## Support
If you have questions, concerns, bug reports, etc. See [CONTRIBUTING](CONTRIBUTING.md).

## License
http-problemdetails-go is Open Source software released under the [Apache 2.0 license](https://www.apache.org/licenses/LICENSE-2.0.html).

## Original Contributors
- Steve Sefton, Evernorth
- Shellee Stewart, Evernorth

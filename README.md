# http-problemdetails-go

An implementation of [IETF RFC 9457 Problem Details for HTTP APIs](https://www.rfc-editor.org/rfc/rfc9457.html) which is a specification for a standard error structure for HTTP APIs.


## How it works

## Installation
```go get -u github.com/Evernorth/http-problemdetails-go```

## Features


## Usage
### Creating a simple HTTP ProblemDetails
You can create a ProblemDetails instance by simply calling functions like `problemdetails.NewInternalServerError()`, `problemdetails.NewNotFound()` and `problemdetails.NewBadRequest()`.
```go
package employee

import (
    "github.com/Evernorth/common-go/http/problemdetails"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    "net/http"
)

func getEmployee(httpRespWriter http.ResponseWriter, httpReq *http.Request) {
    
    err := doSomething()
    if err != nil {
        httpRespWriter.WriteHeader(http.StatusInternalServerError)
        render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError())
        return
    }
    ...
}
```
### Creating a complex HTTP ProblemDetails
You can easily override the default ProblemDetails title and detail fields and provide custom extension fields by using the `WithTitle`, `WithDetail` and `WithExtension` functions.
```go
package employee

import (
    "github.com/Evernorth/common-go/http/problemdetails"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    "net/http"
)

func getEmployee(httpRespWriter http.ResponseWriter, httpReq *http.Request) {
    
    err := doSomething()
    if err != nil {
        httpRespWriter.WriteHeader(http.StatusInternalServerError)
        render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError().WithTitle(
            "KABOOM!!!").WithDetail("The unthinkable has occurred.").WithExtension(
                "example1", "test").WithExtension(
                    "example2", "this could be a struct if you like"))
        return
    }
    ...
}
```

## Dependencies
See the [go.mod](go.mod) file.

## Support
If you have questions, concerns, bug reports, etc. See [CONTRIBUTING](CONTRIBUTING.md).

## License
http-problemdetails-go is open source software released under the [Apache 2.0 license](https://www.apache.org/licenses/LICENSE-2.0.html).

## Original Contributors
- Steve Sefton, Evernorth
- Shellee Stewart, Evernorth
- Neil Powell, Evernorth
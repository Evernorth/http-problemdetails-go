# http-problemdetails-go

An implementation of [IETF RFC 9457 Problem Details for HTTP APIs](https://www.rfc-editor.org/rfc/rfc9457.html) which is a specification for a standard error structure for HTTP APIs.


## Dependencies
See the [go.mod](go.mod) file.

## Building from Source

Detailed instructions on how to build the project from source. Also note where to get pre-built distribution, if building is not required.

## Installation
```go get -u github.com/Evernorth/http-problemdetails-go```

## Configuration

If the software is configurable, describe it in detail, either here or in other documentation to which you link.


## Usage
### Creating a simple HTTP ProblemDetails
You can create a ProblemDetails instance by simply calling functions like `problemdetails.NewInternalServerError()`, `problemdetails.NewNotFound()` and `problemdetails.NewBadRequest()`.
```go
package employee

import (
    "github.com/Evernorth/http-problemdetails-go/http/problemdetails"
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
    "github.com/Evernorth/http-problemdetails-go/http/problemdetails"
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
## How to test the software

If the software includes automated tests, detail how to run those tests.

## Known issues

Document any known significant shortcomings with the software.

## Getting help

If you have questions, concerns, bug reports, etc. See [CONTRIBUTING](CONTRIBUTING.md).

## Getting involved

This section should detail why people should get involved and describe key areas you are
currently focusing on; e.g., trying to get feedback on features, fixing certain bugs, building
important pieces, etc.

General instructions on _how_ to contribute should be stated with a link to [CONTRIBUTING](CONTRIBUTING.md).

## License
http-problemdetails-go is Open Source software released under the [Apache 2.0 license](https://www.apache.org/licenses/LICENSE-2.0.html).

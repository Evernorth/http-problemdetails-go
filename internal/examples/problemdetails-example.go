package main

import (
	"fmt"
	"github.com/Evernorth/http-problemdetails-go/http/problemdetails"
)

func main() {
	// Example of creating a ProblemDetails instance for HTTP 500 Internal Server Error
	internalServerError := problemdetails.NewInternalServerError()
	fmt.Printf("Type: %s, Status: %d, Title: %s\n", internalServerError.Type, internalServerError.Status, internalServerError.Title)

	// Example of creating a ProblemDetails instance for HTTP 400 Bad Request
	badRequest := problemdetails.NewBadRequest().WithTitle("This is bad").WithDetail("This is a very bad request")
	fmt.Printf("Type: %s, Status: %d, Title: %s, Detail: %s\n", badRequest.Type, badRequest.Status, badRequest.Title, badRequest.Detail)

	// Example of creating a ProblemDetails instance for HTTP 404 Not Found
	notFound := problemdetails.NewNotFound()
	fmt.Printf("Type: %s, Status: %d, Title: %s\n", notFound.Type, notFound.Status, notFound.Title)

	// Example of creating a ProblemDetails instance for HTTP 418 I'm a teapot
	teapot := problemdetails.NewTeapot().WithDetail(problemdetails.Teapot.Detail)
	fmt.Printf("Type: %s, Status: %d, Title: %s\n", teapot.Type, teapot.Status, teapot.Title)

}

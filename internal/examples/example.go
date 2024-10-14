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
		httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
		httpRespWriter.WriteHeader(http.StatusMethodNotAllowed)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewMethodNotAllowed())
		return
	}

	// Get the example-type from the query parameters
	queryParams := httpReq.URL.Query()
	exampleType := queryParams.Get("example-type")

	if exampleType == "basic" {

		// Return a 500 Internal Server Error, using the NewInternalServerError function.
		httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
		httpRespWriter.WriteHeader(http.StatusInternalServerError)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError())

	} else if exampleType == "advanced" {

		// Return a 500 Internal Server Error, using the NewInternalServerError and With* functions.
		httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
		httpRespWriter.WriteHeader(http.StatusInternalServerError)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError().
			WithTitle("KABOOM!!!").
			WithDetail("The unthinkable has occurred.").
			WithExtension("example1", "test").
			WithExtension("example2", "this could be a struct if you like"))

	} else {

		// Return a 400 Bad Request, using the NewBadRequest function.
		httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
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

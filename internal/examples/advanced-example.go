package main

import (
	"errors"
	"github.com/Evernorth/problemdetails"
	"net/http"
)

func getSomethingAdvanced(httpRespWriter http.ResponseWriter, httpReq *http.Request) {

	err := doSomethingAdvanced()
	if err != nil {
		httpRespWriter.WriteHeader(http.StatusInternalServerError)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError().WithTitle("KABOOM!!!").WithDetail("The unthinkable has occurred.").WithExtension(
			"example1", "test").WithExtension(
			"example2", "this could be a struct if you like"))
		httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)

		return
	}
}

func doSomethingAdvanced() error {
	return errors.New("look who stepped in the room")
}

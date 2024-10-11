package main

import (
	"errors"
	"github.com/Evernorth/problemdetails"
	"github.com/go-chi/render"
	"net/http"
)

func getSomethingBasic(httpRespWriter http.ResponseWriter, httpReq *http.Request) {

	err := doSomethingBasic()
	if err != nil {
		httpRespWriter.WriteHeader(http.StatusInternalServerError)
		render.JSON(httpRespWriter, httpReq, problemdetails.NewInternalServerError())
		httpRespWriter.Header().Set("Content-Type", problemdetails.MimeType)
		return
	}
}

func doSomethingBasic() error {
	return errors.New("look who stepped in the room")
}

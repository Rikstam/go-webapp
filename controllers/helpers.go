package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, destination interface{}) error {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(destination, r.PostForm); err != nil {
		return err
	}
	return nil
}

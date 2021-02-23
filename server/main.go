// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}

	http.ListenAndServe(":"+port, &app.Handler{
		Title: "Dicetalk",
		Styles: []string{
			"/web/main.css",
			"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css",
		},
		Scripts: []string{
			"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js",
		},
	})
}

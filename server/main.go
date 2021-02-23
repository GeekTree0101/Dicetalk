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
		},
	})
}

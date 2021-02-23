package main

import (
	"github.com/GeekTree0101/Dicetalk.git/view"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	app.Route("/", &view.DicetalkView{})
	app.Route("/register", &view.RegisterView{})
	app.Run()
}

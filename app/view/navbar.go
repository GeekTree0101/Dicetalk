package view

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// NavbarItem : navigation bar item
type NavbarItem struct {
	Text string
	Herf string
}

// Navbar : navigation bar
func Navbar(title string, items []NavbarItem) app.UI {
	return app.Nav().Body(
		app.Div().Class("nav-wrapper grey darken-4").Body(
			app.A().Class("logo").Text(title),
			app.Ul().Class("right").Body(
				app.Range(items).Slice(func(i int) app.UI {
					return app.Li().Body(
						app.A().Href(items[i].Herf).Text(items[i].Text),
					)
				}),
			),
		),
	)
}

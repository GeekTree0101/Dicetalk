package view

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// RegisterView : register view
type RegisterView struct {
	app.Compo
}

// Render : render view
func (v *RegisterView) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.H1().Body(
				app.Text("Register"),
			),
		),
	)
}

// DidTapRegisterButton : did tap register button action
func (v *RegisterView) DidTapRegisterButton(ctx app.Context, e app.Event) {

	// TODO: pick
	v.Update()
}

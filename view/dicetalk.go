package view

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// DicetalkView : dice talk view
type DicetalkView struct {
	app.Compo
	name string
}

// Render : render view
func (v *DicetalkView) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.H1().Body(
				app.Text("Dice view"),
			),
		),
	)
}

// DidTapDiceButton : did tap dice button action
func (v *DicetalkView) DidTapDiceButton(ctx app.Context, e app.Event) {

	// TODO: pick
	v.Update()
}

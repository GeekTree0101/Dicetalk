package view

import (
	"github.com/GeekTree0101/Dicetalk/app/service"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// DicetalkView : dice talk view
type DicetalkView struct {
	app.Compo
	name          string
	TopicService  *service.TopicService
	MemberService *service.MemberService
	currentTopic  string
	currentMember string
}

// Render : render view
func (v *DicetalkView) Render() app.UI {
	return app.Div().Body(
		Navbar("Dicetalk", []NavbarItem{
			{
				Herf: "/register",
				Text: "환경설정",
			},
		}),
		v.Content(),
	)
}

// Content : content view
func (v *DicetalkView) Content() app.UI {

	topics := v.TopicService.GetAll()
	members := v.MemberService.GetAll()

	if len(topics) > 0 && len(members) > 0 {

		return app.Div().Class("content").Body(
			app.If(v.isValidContent(), v.TopicView()),
			app.Button().Class("btn-large waves-effect waves-light").Text("🎲 주사위 굴리기").OnClick(v.DidTapDiceButton),
		)
	}

	return app.Div().Class("content").Body(
		app.Div().Body(
			app.P().Class("error-title").Text("🤞 상단바에 있는 환경설정을 눌려서\n맴버와 주제를 잔뜩추가해주세요."),
		),
	)
}

// TopicView : topic view
func (v *DicetalkView) TopicView() app.UI {
	return app.Div().Class("topic-container").Body(
		app.P().Class("topic row").Text("\""+v.currentTopic+"\""),
		app.P().Class("topic row").Text(v.currentMember),
	)
}

func (v *DicetalkView) isValidContent() bool {

	if len(v.currentMember) > 0 && len(v.currentTopic) > 0 {
		return true
	}

	return false
}

// DidTapDiceButton : did tap dice button action
func (v *DicetalkView) DidTapDiceButton(ctx app.Context, e app.Event) {

	v.currentMember = v.MemberService.GetMember()
	v.currentTopic = v.TopicService.GetTopic()

	v.Update()
}

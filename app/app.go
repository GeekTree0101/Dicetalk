package main

import (
	"github.com/GeekTree0101/Dicetalk/app/service"
	"github.com/GeekTree0101/Dicetalk/app/view"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	ts := service.TopicService{}
	ms := service.MemberService{}

	ts.Appends([]string{"david kill", "sonic boom"})
	ms.Appends([]string{"david", "martry", "eric"})

	app.Route("/", &view.DicetalkView{
		TopicService:  &ts,
		MemberService: &ms,
	})

	app.Route("/register", &view.RegisterView{
		TopicService:  &ts,
		MemberService: &ms,
	})

	app.Run()
}

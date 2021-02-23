package main

import (
	"github.com/GeekTree0101/Dicetalk/service"
	"github.com/GeekTree0101/Dicetalk/view"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	ts := service.TopicService{}
	ms := service.MemberService{}

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

package view

import (
	"github.com/GeekTree0101/Dicetalk/app/service"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// RegisterView : register view
type RegisterView struct {
	app.Compo
	TopicService  *service.TopicService
	MemberService *service.MemberService
	currentTopic  string
	currentMember string
}

// ViewType : view type
type ViewType string

const (
	// TopicViewType : topic view type
	TopicViewType ViewType = "topic"
	// MemberViewType : member view type
	MemberViewType ViewType = "member"
)

func (t ViewType) placeholder() string {
	switch t {
	case TopicViewType:
		return "주제를 입력해주세요"
	case MemberViewType:
		return "맴버 이름을 입력해주세요"
	default:
		return ""
	}
}

func (t ViewType) title() string {
	switch t {
	case TopicViewType:
		return "주제 선정하기"
	case MemberViewType:
		return "맴버 선정하기"
	default:
		return ""
	}
}

// Render : render view
func (v *RegisterView) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("navbar").Body(
			app.P().Class("navbar-title").Text("Dicetalk 설정하기"),
		),
		v.Section(MemberViewType, v.MemberService.GetAll()),
		v.Section(TopicViewType, v.TopicService.GetAll()),
	)
}

// Section : section view
func (v *RegisterView) Section(viewType ViewType, items []string) app.UI {
	return app.Div().Class("item-section").Body(
		app.H3().Text(viewType.title()),
		v.Field(viewType),
		app.Div().Class("item-container").Body(
			app.Range(items).Slice(func(i int) app.UI {
				return app.Button().
					Class("item-tag").
					Text(items[i]).
					Value(items[i]).
					OnClick(func() app.EventHandler {
						switch viewType {
						case TopicViewType:
							return v.DidTapDeleteTopic
						case MemberViewType:
							return v.DidTapDeleteMember
						}
						return nil
					}())
			}),
		),
	)
}

// Field : field
func (v *RegisterView) Field(viewType ViewType) app.UI {
	return app.Div().Class("in-line").Body(
		app.Input().
			Class("text-input").
			Type("text").
			Name("name").
			Value("").
			Placeholder(viewType.placeholder()).
			OnChange(func() app.EventHandler {
				switch viewType {
				case TopicViewType:
					return v.OnChangeTopicField
				case MemberViewType:
					return v.OnChangeMemberField
				}
				return nil
			}()),
		app.Input().
			Class("input-button").
			Type("button").
			Name("name").
			Value("추가").
			OnClick(func() app.EventHandler {
				switch viewType {
				case TopicViewType:
					return v.DidTapRegisterTopicButton
				case MemberViewType:
					return v.DidTapRegisterMemberButton
				}
				return nil
			}()),
	)
}

// OnChangeMemberField : member field on change
func (v *RegisterView) OnChangeMemberField(ctx app.Context, e app.Event) {

	v.currentMember = ctx.JSSrc.Get("value").String()
}

// OnChangeTopicField : topic field on change
func (v *RegisterView) OnChangeTopicField(ctx app.Context, e app.Event) {

	v.currentTopic = ctx.JSSrc.Get("value").String()
}

// DidTapRegisterMemberButton : did tap register member button action
func (v *RegisterView) DidTapRegisterMemberButton(ctx app.Context, e app.Event) {

	if len(v.currentMember) > 0 {
		v.MemberService.Append(v.currentMember)
	}

	v.Update()
}

// DidTapRegisterTopicButton : did tap register topic button action
func (v *RegisterView) DidTapRegisterTopicButton(ctx app.Context, e app.Event) {

	if len(v.currentTopic) > 0 {
		v.TopicService.Append(v.currentTopic)
	}

	v.Update()
}

// DidTapDeleteMember : did tap delete member
func (v *RegisterView) DidTapDeleteMember(ctx app.Context, e app.Event) {

	m := ctx.JSSrc.Get("value").String()
	v.MemberService.Delete(m)
	v.Update()
}

// DidTapDeleteTopic : did tap delete topic
func (v *RegisterView) DidTapDeleteTopic(ctx app.Context, e app.Event) {

	t := ctx.JSSrc.Get("value").String()
	v.TopicService.Delete(t)
	v.Update()
}

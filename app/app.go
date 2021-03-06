package main

import (
	"github.com/GeekTree0101/Dicetalk/app/service"
	"github.com/GeekTree0101/Dicetalk/app/view"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	ts := service.TopicService{}
	ms := service.MemberService{}

	ms.Appends([]string{"마티", "데이빗", "다니엘", "레이", "로니", "헤세드", "레오"})

	ts.Appends([]string{
		"첫사랑이야기 해주세요",
		"가장 기억에 남는 여행지 이야기해주세요",
		"어떠한 직업도 가질 수 있을 때, 어떤 일을 하시겠나요? (현재 당근마켓 월급과 동일한 월급을 받음. 현재 하고 있는 직무 제외)",
		"작년에 찍었던 사진 중 가장 인상 깊은 사진은 무엇인가요? (직접 찍은 사진 중 마음에 드는 것이 없다면, 스크린 샷한 이미지도 좋아요)",
		"최근 3달 안에 보았던 책(or)영화 공유하기!4. 아무도 시키지 않았으며, 금전적인 보상도 없는 일 중에서 자발적으로 하게 되는 일이 있나요?",
		"자신의 2020년을 한 문장으로 정리해보면은 어떻게 표현할 수 있을까요?",
		"자신이 재능 있다고 생각하는 것과 재능이 하나도 없지만 잘하고 싶은 건 무엇인가요? (일과 상관없어도 좋아요!)",
		"죽기전에 꼭 한번 해보고 싶은 건 무엇인가요?",
		"당신이 한 가장 불법적 인 일은 무엇입니까?",
		"당신이 누군가에게 준 최고의 선물은 무엇이었습니까?",
		"어떤 음식을 절대 시도하지 않겠습니까?",
		"당신의 완벽한 버거는 무엇입니까?",
		"‘사랑’이 뭔지 설명해 주 시겠어요?",
		"지금까지 맛본것중에서 최악의 것은 무엇었나요?",
		"집에 불이 났을 때 무엇을 챙기겠습니까?",
		"완벽한 아침 식사는 무엇입니까?",
		"당신은 아직도 어떤 유치한 것을 좋아합니까?",
		"회사에서 가장 창피했던 순간은 언제 였나요?",
		"장님이 되시지만 항상 머리가 꽉 찬 상태가 되시겠습니까, 아니면 대머리가 되어도 시력을 잃지 않겠습니까?",
		"귀하의 의견으로는 가장 쓸모없는 동물은 무엇입니까?",
		"회사에서 사용하는 당신의 이름은 무엇을 의미합니까?",
		"Swift보다 Kotlin이 나쁜점에 대해서 예기해주세요",
		"Kotlin보다 Swift의 나쁜점에 대해서 예기해주세요",
		"코로나가 끝난다면 가장 가고 싶은 여행지를 얘기해주세요.",
		"여행지에서 기억에 남는 액티비티를 얘기해주세요.",
		"추천해주고 싶은 여행지가 있다면?",
		"개발자가 된 계기를 얘기해주세요.",
		"개발 외에(직업과 상관없이) 배우고 싶은 게 있다면?",
		"내가 남들보다 개발을 압도적으로 잘하는 비결은?",
		"내 개인만의 꿀 피부 관리 비결은?",
		"반드시 목디스크와 허리디스크중 하나를 고른다면 어떤 병을 고르시겠나요? 그 이유는?",
	})

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

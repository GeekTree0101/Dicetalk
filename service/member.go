package service

// MemberService : member service
type MemberService struct {
	members []string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Appends : append members
func (s *MemberService) Appends(members []string) {
	s.members = members
}

// GetMember : get random member
func (s *TopicService) GetMember() string {
	i := rand.Intn(len(s.members))
	return members[i]
}
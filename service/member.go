package service

import (
	"math/rand"
	"time"
)

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
func (s *MemberService) GetMember() string {
	i := rand.Intn(len(s.members))
	return s.members[i]
}

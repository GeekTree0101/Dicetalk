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

// Append : append member
func (s *MemberService) Append(member string) {
	s.members = append(s.members, member)
}

// GetAll : get all members
func (s *MemberService) GetAll() []string {
	return s.members
}

// GetMember : get random member
func (s *MemberService) GetMember() string {
	i := rand.Intn(len(s.members))
	return s.members[i]
}

// Delete : delete topic
func (s *MemberService) Delete(m string) {
	s.members = removeMember(s.members, m)
}

func removeMember(ms1 []string, m1 string) []string {

	ms2 := []string{}

	for _, m2 := range ms1 {
		if m2 != m1 {
			ms2 = append(ms2, m2)
		}
	}

	return ms2
}

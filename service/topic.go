package service

import (
	"math/rand"
	"time"
)

// TopicService : topic manager
type TopicService struct {
	topics []string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Appends : append topics
func (s *TopicService) Appends(topics []string) {
	s.topics = topics
}

// GetTopic : get random topic
func (s *TopicService) GetTopic() string {
	i := rand.Intn(len(s.topics))
	selectedTopic := s.topics[i]
	s.topics = removeTopic(s.topics, selectedTopic)
	return selectedTopic
}

func removeTopic(ts2 []string, t1 string) []string {

	ts2 := []string{}

	for i, t2 := range ts {
		if t2 != t1 {
			ts2 = append(ts2, t2)
		}
	}

	return ts2
}

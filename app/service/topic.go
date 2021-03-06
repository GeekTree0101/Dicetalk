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

// Append : append topic
func (s *TopicService) Append(topic string) {
	s.topics = append(s.topics, topic)
}

// GetAll : get all members
func (s *TopicService) GetAll() []string {
	return s.topics
}

// GetTopic : get random topic
func (s *TopicService) GetTopic() string {
	i := rand.Intn(len(s.topics))
	return s.topics[i]
}

// Delete : delete topic
func (s *TopicService) Delete(t string) {
	s.topics = removeTopic(s.topics, t)
}

func removeTopic(ts1 []string, t1 string) []string {

	ts2 := []string{}

	for _, t2 := range ts1 {
		if t2 != t1 {
			ts2 = append(ts2, t2)
		}
	}

	return ts2
}

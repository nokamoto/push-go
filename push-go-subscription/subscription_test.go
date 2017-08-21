package main

import (
	"testing"
	"github.com/nokamoto/push-go/proto"
)

func TestSubscription_Subscribe(t *testing.T) {
	s := NewSubscription()

	entry := &push.Subscription{
		Topic: &push.Topic{Name: "x"},
		Key: &push.Id{Id: "y"},
	}

	t.Logf("subscribe new entry: %v", entry)
	if _, err := s.Subscribe(nil, entry); err != nil {
		t.Error(err.Error())
	}

	if ids, ok := s.topics[*entry.Topic]; !ok {
		t.Errorf("%v not found: %v", entry, s.topics)
	} else if len(ids) != 1 {
		t.Errorf("len(%v) is not 1", ids)
	} else if ids[0].Id != entry.Key.Id {
		t.Errorf("%v is not equal to %v", ids[0], entry.Key)
	}
}

func TestSubscription_Unsubscribe(t *testing.T) {
	s := NewSubscription()

	entry := &push.Subscription{
		Topic: &push.Topic{Name: "x"},
		Key: &push.Id{Id: "y"},
	}

	t.Logf("unsubscribe undefined entry: %v", entry)
	if _, err := s.Unsubscribe(nil, entry); err != nil {
		t.Error(err.Error())
	}

	t.Logf("subscribe new entry: %v", entry)
	if _, err := s.Subscribe(nil, entry); err != nil {
		t.Error(err.Error())
	}

	t.Logf("unsubscribe defined entry: %v", entry)
	if _, err := s.Unsubscribe(nil, entry); err != nil {
		t.Error(err.Error())
	}

	if ids, ok := s.topics[*entry.Topic]; !ok && len(ids) != 0 {
		t.Errorf("%v is not empty", s.topics)
	}
}

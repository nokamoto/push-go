package main

import (
	"testing"
	"github.com/nokamoto/push-go/proto"
)

func TestLogging_Info(t *testing.T) {
	l := NewLog()

	entry := push.Log{Text: "info text"}

	t.Logf("info: %v", entry)
	if _, err := l.Info(nil, &entry); err != nil {
		t.Error(err.Error())
	}

	if size := len(l.info); size != 1 {
		t.Errorf("expected 1 text but actual %d: %v", size, l.info)
	}

	if l.info[0].Text != entry.Text {
		t.Errorf("expected %v but actual %v", entry, l.info[0])
	}
}

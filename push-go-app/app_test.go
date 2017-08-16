package main

import (
	"testing"
	"github.com/nokamoto/push-go/proto"
	"reflect"
)

func TestApp_GetFirebaseCloudMessaging(t *testing.T) {
	a := new(App)

	t.Log("get uninitialized firebase expected")
	if empty, err := a.GetFirebaseCloudMessaging(nil, nil); err != nil {
		t.Error(err.Error())
	} else if d := new(push.FirebaseCloudMessagingApp); !reflect.DeepEqual(empty, d) {
		t.Errorf("%v != %v", *empty, *d)
	}

	expected := push.FirebaseCloudMessagingApp{ApiKey: "default"}
	t.Logf("set %v", expected)
	if _, err := a.SetFirebaseCloudMessaging(nil, &expected); err != nil {
		t.Error(err.Error())
	}

	if actual, err := a.GetFirebaseCloudMessaging(nil, nil); err != nil {
		t.Error(err.Error())
	} else if !reflect.DeepEqual(actual, &expected) {
		t.Errorf("%v != %v", *actual, expected)
	}
}

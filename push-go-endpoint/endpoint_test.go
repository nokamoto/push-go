package main

import (
	"testing"
	"github.com/nokamoto/push-go/proto"
	"reflect"
)

func TestEndpoint_Set(t *testing.T) {
	e := NewEndpoint()

	x := push.Id{Id: "x"}

	token := push.FirebaseCloudMessagingEndpoint{Token: "y"}

	entry := push.SetFirebaseCloudMessagingEndpoint{
		Key: &x,
		Value: &token,
	}

	t.Logf("set %v", entry)
	if _, err := e.Set(nil, &entry); err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(e.endpoints[x], &token) {
		t.Errorf("%v != %v", *e.endpoints[x], token)
	}
}

func TestEndpoint_Delete(t *testing.T) {
	e := NewEndpoint()

	x := push.Id{Id: "x"}

	token := push.FirebaseCloudMessagingEndpoint{Token: "y"}

	entry1 := push.SetFirebaseCloudMessagingEndpoint{
		Key: &x,
		Value: &token,
	}

	t.Logf("set %v", entry1)
	e.Set(nil, &entry1)


	entry2 := push.DeleteFirebaseCloudMessagingEndpoint{Value: &token}

	t.Logf("delete %v", entry2)
	if _, err := e.Delete(nil, &entry2); err != nil {
		t.Error(err.Error())
	}

	if v, exists := e.endpoints[x]; exists {
		t.Errorf("expected not exist but actual %v", v)
	}

	t.Logf("delete %v again", entry2)
	if _, err := e.Delete(nil, &entry2); err != nil {
		t.Error(err.Error())
	}
}

func TestEndpoint_Update(t *testing.T) {
	e := NewEndpoint()

	x := push.Id{Id: "x"}

	token := push.FirebaseCloudMessagingEndpoint{Token: "y"}

	entry1 := push.SetFirebaseCloudMessagingEndpoint{
		Key: &x,
		Value: &token,
	}

	if _, err := e.Set(nil, &entry1); err != nil {
		t.Error(err.Error())
	}

	newToken := push.FirebaseCloudMessagingEndpoint{Token: "z"}

	update := push.UpdateFirebaseCloudMessagingEndpoint{
		OldValue: &token,
		NewValue: &newToken,
	}

	t.Logf("update %v", update)
	if _, err := e.Update(nil, &update); err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(e.endpoints[x], &newToken) {
		t.Errorf("%v != %v", *e.endpoints[x], newToken)
	}
}

package main

import (
	"github.com/nokamoto/push-go/proto"
)

type EndpointDeleteTopic struct {}

func (e EndpointDeleteTopic)Name() []string {
	return []string{"topic"}
}

func (e EndpointDeleteTopic)Run(opts Options, args []string) error {
	return delete(opts, args, func(v string) *push.FirebaseCloudMessagingEndpoint {
		return &push.FirebaseCloudMessagingEndpoint{Topic: v}
	})
}


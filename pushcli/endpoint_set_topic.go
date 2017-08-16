package main

import (
	"github.com/nokamoto/push-go/proto"
)

type EndpointSetTopic struct {}

func (e EndpointSetTopic)Name() []string {
	return []string{"topic"}
}

func (e EndpointSetTopic)Run(opts Options, args []string) error {
	return set(opts, args, func(v string) *push.FirebaseCloudMessagingEndpoint {
		return &push.FirebaseCloudMessagingEndpoint{Topic: v}
	})
}

package main

import (
	"github.com/nokamoto/push-go/proto"
)

type EndpointSetToken struct {}

func (e EndpointSetToken)Name() []string {
	return []string{"token"}
}

func (e EndpointSetToken)Run(opts Options, args []string) error {
	return set(opts, args, func(v string) *push.FirebaseCloudMessagingEndpoint {
		return &push.FirebaseCloudMessagingEndpoint{Token: v}
	})
}


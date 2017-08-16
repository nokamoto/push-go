package main

import (
	"github.com/nokamoto/push-go/proto"
)

type EndpointDeleteToken struct {}

func (e EndpointDeleteToken)Name() []string {
	return []string{"token"}
}

func (e EndpointDeleteToken)Run(opts push.Options, args []string) error {
	return delete(opts, args, func(v string) *push.FirebaseCloudMessagingEndpoint {
		return &push.FirebaseCloudMessagingEndpoint{Token: v}
	})
}


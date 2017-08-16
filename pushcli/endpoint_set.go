package main

import (
	"errors"
	"github.com/nokamoto/push-go/proto"
	"fmt"
	"golang.org/x/net/context"
)

type EndpointSet struct {}

func (e EndpointSet)Name() []string {
	return []string{"set", "s"}
}

func (e EndpointSet)Run(opts push.Options, args []string) error {
	if len(args) != 2 {
		return errors.New("Usage: set id token")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingEndpointServiceClient(conn)
	req := &push.SetFirebaseCloudMessagingEndpoint{
		Key: &push.Id{Id: args[0]},
		Value: &push.FirebaseCloudMessagingEndpoint{Token: args[1]},
	}

	fmt.Printf("set token: %v\n", req)
	_, err = client.Set(context.Background(), req)

	return err
}

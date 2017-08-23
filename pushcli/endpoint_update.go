package main

import (
	"errors"
	"github.com/nokamoto/push-go/proto"
	"fmt"
	"golang.org/x/net/context"
)

type EndpointUpdate struct {}

func (e EndpointUpdate)Name() []string {
	return []string{"update", "u"}
}

func (e EndpointUpdate)Run(opts push.Options, args []string) error {
	if len(args) != 2 {
		return errors.New("Usage: update old-token new-token")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingEndpointServiceClient(conn)
	req := &push.UpdateFirebaseCloudMessagingEndpoint{
		OldValue: &push.FirebaseCloudMessagingEndpoint{Token: args[0]},
		NewValue: &push.FirebaseCloudMessagingEndpoint{Token: args[1]},
	}

	fmt.Printf("update token: %v\n", req)
	_, err = client.Update(context.Background(), req)

	return err
}


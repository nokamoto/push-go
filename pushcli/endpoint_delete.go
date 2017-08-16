package main

import (
	"errors"
	"github.com/nokamoto/push-go/proto"
	"fmt"
	"golang.org/x/net/context"
)

type EndpointDelete struct {}

func (e EndpointDelete)Name() []string {
	return []string{"delete", "d"}
}

func (e EndpointDelete)Run(opts push.Options, args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: delete token")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingEndpointServiceClient(conn)
	req := &push.DeleteFirebaseCloudMessagingEndpoint{
		Value: &push.FirebaseCloudMessagingEndpoint{Token: args[0]},
	}

	fmt.Printf("delete token: %v\n", req)
	_, err = client.Delete(context.Background(), req)

	return err
}


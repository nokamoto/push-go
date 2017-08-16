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

func (e EndpointDelete)SubCommands() []SubCommand {
	return []SubCommand{EndpointDeleteToken{}, EndpointDeleteTopic{}}
}

func (e EndpointDelete)Run(opts Options, args []string) error {
	return run(e, opts, args)
}

func delete(opts Options, args []string, f func(string)*push.FirebaseCloudMessagingEndpoint) error {
	if len(args) != 1 {
		return errors.New("Usage: type value")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingEndpointServiceClient(conn)
	req := &push.DeleteFirebaseCloudMessagingEndpoint{
		Value: f(args[0]),
	}

	fmt.Printf("delete token: %v\n", req)
	_, err = client.Delete(context.Background(), req)

	return err
}



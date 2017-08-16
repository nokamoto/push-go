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

func (e EndpointSet)SubCommands() []SubCommand {
	return []SubCommand{EndpointSetToken{}, EndpointSetTopic{}}
}

func (e EndpointSet)Run(opts Options, args []string) error {
	return run(e, opts, args)
}

func set(opts Options, args []string, f func(string)*push.FirebaseCloudMessagingEndpoint) error {
	if len(args) != 2 {
		return errors.New("Usage: type id value")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingEndpointServiceClient(conn)
	req := &push.SetFirebaseCloudMessagingEndpoint{
		Key: &push.Id{Id: args[0]},
		Value: f(args[1]),
	}

	fmt.Printf("set token: %v\n", req)
	_, err = client.Set(context.Background(), req)

	return err
}


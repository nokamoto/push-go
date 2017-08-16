package main

import (
	"errors"
	"github.com/nokamoto/push-go/proto"
	"fmt"
	"golang.org/x/net/context"
	"io"
)

type EndpointGet struct {}

func (e EndpointGet)Name() []string {
	return []string{"get", "g"}
}

func (e EndpointGet)Run(opts push.Options, args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: get id")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingEndpointServiceClient(conn)
	req := &push.Id{Id: args[0]}

	if stream, err := client.Get(context.Background(), req); err != nil {
		return err
	} else {
		for {
			value, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			fmt.Printf("%v\n", value)
		}
	}

	return nil
}


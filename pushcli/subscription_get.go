package main

import (
	"github.com/nokamoto/push-go/proto"
	"errors"
	"io"
	"fmt"
	"golang.org/x/net/context"
	"github.com/golang/protobuf/jsonpb"
)

type SubscriptionGet struct {}

func (a SubscriptionGet)Name() []string {
	return []string{"get", "g"}
}

func (a SubscriptionGet)Run(opts push.Options, args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: get topic")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewSubscriptionServiceClient(conn)
	req := &push.Topic{Name: args[0]}

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

			json, err := new(jsonpb.Marshaler).MarshalToString(value)
			if err != nil {
				return err

			}
			fmt.Printf("%v\n", json)
		}
	}

	return nil
}

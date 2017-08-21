package main

import (
	"github.com/nokamoto/push-go/proto"
	"errors"
	"fmt"
	"golang.org/x/net/context"
)

type SubscriptionUnsubscribe struct {}

func (a SubscriptionUnsubscribe)Name() []string {
	return []string{"unsubscribe", "u"}
}

func (a SubscriptionUnsubscribe)Run(opts push.Options, args []string) error {
	if len(args) != 2 {
		return errors.New("Usage: unsubscribe topic id")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewSubscriptionServiceClient(conn)
	req := &push.Subscription{
		Topic: &push.Topic{Name: args[0]},
		Key: &push.Id{Id: args[1]},
	}

	fmt.Printf("unsubscribe: %v\n", req)

	_, err = client.Unsubscribe(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

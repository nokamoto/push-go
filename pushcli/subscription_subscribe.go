package main

import (
	"github.com/nokamoto/push-go/proto"
	"errors"
	"fmt"
	"golang.org/x/net/context"
)

type SubscriptionSubscribe struct {}

func (a SubscriptionSubscribe)Name() []string {
	return []string{"subscribe", "s"}
}

func (a SubscriptionSubscribe)Run(opts push.Options, args []string) error {
	if len(args) != 2 {
		return errors.New("Usage: subscribe topic id")
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

	fmt.Printf("subscribe: %v\n", req)

	_, err = client.Subscribe(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

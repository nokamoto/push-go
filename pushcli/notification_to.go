package main

import (
	"github.com/nokamoto/push-go/proto"
	"errors"
	"fmt"
	"context"
)

type NotificationTo struct {
	names []string
	request func(string, string)*push.FirebaseCloudMessagingNotification
}

func NewNotificationTo(names []string, f func(*push.FirebaseCloudMessagingNotification, string)) NotificationTo {
	return 	NotificationTo{
		names: names,
		request: func(x string, json string)*push.FirebaseCloudMessagingNotification {
			n := &push.FirebaseCloudMessagingNotification{
				Payload: &push.FirebaseCloudMessagingNotification_Payload{
					Json: json,
				},
			}
			f(n, x)
			return n
		},
	}
}

func (n NotificationTo)Name() []string {
	return n.names
}

func (n NotificationTo)Run(opts push.Options, args []string) error {
	if len(args) != 2 {
		return errors.New(fmt.Sprintf("Usage: %v value json", n.Name()[0]))
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewFirebaseCloudMessagingServiceClient(conn)
	req := n.request(args[0], args[1])

	fmt.Printf("send: %v", req)

	_, err = client.Send(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

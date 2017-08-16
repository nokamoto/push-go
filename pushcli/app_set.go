package main

import (
	"errors"
	"github.com/nokamoto/push-go/proto"
	"fmt"
	"golang.org/x/net/context"
)

type AppSet struct {}

func (a AppSet)Name() []string {
	return []string{"set", "s"}
}

func (a AppSet)Run(opts Options, args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: set api-key")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewAppServiceClient(conn)
	req := &push.FirebaseCloudMessagingApp{ApiKey: args[0]}

	fmt.Printf("set: api-key=%s\n", args[0])

	_, err = client.SetFirebaseCloudMessaging(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

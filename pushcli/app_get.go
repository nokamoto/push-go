package main

import (
	"github.com/nokamoto/push-go/proto"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"fmt"
	"errors"
)

type AppGet struct {}

func (a AppGet)Name() []string {
	return []string{"get", "g"}
}

func (a AppGet)Run(opts Options, args []string) error {
	if len(args) != 0 {
		return errors.New("Usage: get")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewAppServiceClient(conn)
	req := &google_protobuf.Empty{}

	res, err := client.GetFirebaseCloudMessaging(context.Background(), req)
	if err != nil {
		return err
	}

	fmt.Printf("%v", res)

	return nil
}

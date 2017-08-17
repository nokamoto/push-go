package main

import (
	"github.com/nokamoto/push-go/proto"
	"errors"
	"fmt"
	"golang.org/x/net/context"
)

type LogInfo struct {}

func (a LogInfo)Name() []string {
	return []string{"info", "i"}
}

func (a LogInfo)Run(opts push.Options, args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: info text")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewLogServiceClient(conn)
	req := &push.Log{Text: args[0]}

	fmt.Printf("info: %v\n", req)
	_, err = client.Info(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}


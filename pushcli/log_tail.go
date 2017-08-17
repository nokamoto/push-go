package main

import (
	"github.com/nokamoto/push-go/proto"
	"errors"
	"fmt"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"io"
)

type LogTail struct {}

func (a LogTail)Name() []string {
	return []string{"tail", "t"}
}

func (a LogTail)Run(opts push.Options, args []string) error {
	if len(args) != 0 {
		return errors.New("Usage: tail")
	}

	conn, err := opts.Dial()
	if err != nil {
		return err
	}

	defer conn.Close()

	client := push.NewLogServiceClient(conn)
	req := &google_protobuf.Empty{}

	stream, err := client.Tail(context.Background(), req)
	if err != nil {
		return err
	}

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

	return nil
}


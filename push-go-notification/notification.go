package main

import (
	"flag"
	"github.com/nokamoto/push-go/proto"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

type Notification struct {}

func (n * Notification)Send(_ context.Context, notification *push.FirebaseCloudMessagingNotification) (*google_protobuf.Empty, error) {
	return nil, nil
}

func main() {
	server := push.NewServer()

	notification := new(Notification)

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterFirebaseCloudMessagingServiceServer(s, notification)
	})
}

package main

import (
	"github.com/nokamoto/push-go/proto"
	"flag"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

type Endpoint struct {}

func (e *Endpoint)Set(_ context.Context, endpoint *push.SetFirebaseCloudMessagingEndpoint) (*google_protobuf.Empty, error) {
	return nil, nil
}

func (e *Endpoint)Delete(_ context.Context, endpoint *push.DeleteFirebaseCloudMessagingEndpoint) (*google_protobuf.Empty, error) {
	return nil, nil
}

func (e *Endpoint)Get(id *push.Id, stream push.FirebaseCloudMessagingEndpointService_GetServer) error {
	return nil
}

func main() {
	server := push.NewServer()

	endpoint := new(Endpoint)

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterFirebaseCloudMessagingEndpointServiceServer(s, endpoint)
	})
}

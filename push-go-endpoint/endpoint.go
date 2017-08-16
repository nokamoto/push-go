package main

import (
	"github.com/nokamoto/push-go/proto"
	"flag"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"log"
)

type Endpoint struct {
	endpoints map[push.Id]*push.FirebaseCloudMessagingEndpoint
}

func NewEndpoint() *Endpoint {
	return &Endpoint{endpoints: map[push.Id]*push.FirebaseCloudMessagingEndpoint{}}
}

func (e *Endpoint)Set(_ context.Context, endpoint *push.SetFirebaseCloudMessagingEndpoint) (*google_protobuf.Empty, error) {
	log.Printf("set: %v", endpoint)
	e.endpoints[*endpoint.Key] = endpoint.Value
	return new(google_protobuf.Empty), nil
}

func (e *Endpoint)Delete(_ context.Context, endpoint *push.DeleteFirebaseCloudMessagingEndpoint) (*google_protobuf.Empty, error) {
	log.Printf("delete: %v", endpoint)
	for k, v := range e.endpoints {
		if v.Token == endpoint.Value.Token {
			log.Printf("%v deleted", k)
			delete(e.endpoints, k)
		}
	}
	return new(google_protobuf.Empty), nil
}

func (e *Endpoint)Get(id *push.Id, stream push.FirebaseCloudMessagingEndpointService_GetServer) error {
	log.Printf("get: %v", id)
	if x, ok := e.endpoints[*id]; ok {
		if err := stream.Send(x); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	server := push.NewServer()

	endpoint := NewEndpoint()

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterFirebaseCloudMessagingEndpointServiceServer(s, endpoint)
	})
}

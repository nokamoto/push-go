package main

import (
	"github.com/nokamoto/push-go/proto"
	"flag"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"log"
)

type Subscription struct {
	topics map[push.Topic][]*push.Id
}

func NewSubscription() *Subscription {
	return &Subscription{topics: map[push.Topic][]*push.Id{}}
}

func (s *Subscription)Subscribe(_ context.Context, subscription *push.Subscription) (*google_protobuf.Empty, error) {
	log.Printf("subscribe: %v", subscription)

	if ids, ok := s.topics[*subscription.Topic]; ok {
		found := false

		for _, id := range ids {
			found = found || id.Id == subscription.Key.Id
		}

		if !found {
			s.topics[*subscription.Topic] = append(ids, subscription.Key)
		}
	} else {
		s.topics[*subscription.Topic] = []*push.Id{subscription.Key}
	}

	return &google_protobuf.Empty{}, nil
}

func (s *Subscription)Unsubscribe(_ context.Context, subscription *push.Subscription) (*google_protobuf.Empty, error) {
	log.Printf("unsubscribe: %v", subscription)

	if ids, ok := s.topics[*subscription.Topic]; ok {
		index := -1

		for i, id := range ids {
			if id.Id == subscription.Key.Id {
				index = i
			}
		}

		if index != -1 {
			s.topics[*subscription.Topic] = append(ids[:index], ids[index + 1:]...)
		}
	}

	return &google_protobuf.Empty{}, nil
}

func (s *Subscription)Get(topic *push.Topic, stream push.SubscriptionService_GetServer) error {
	log.Printf("get: %v", topic)

	if ids, ok := s.topics[*topic]; ok {
		for _, id := range ids {
			if err := stream.Send(id); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	server := push.NewServer()

	sub := NewSubscription()

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterSubscriptionServiceServer(s, sub)
	})
}

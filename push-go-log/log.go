package main

import (
	"github.com/nokamoto/push-go/proto"
	"flag"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Log struct {
	tail *int

	info []*push.Log
}

func NewLog() *Log {
	return &Log{info: []*push.Log{}}
}

func (l *Log)Info(_ context.Context, info *push.Log) (*empty.Empty, error) {
	log.Printf("info: %v", info)

	info.Timestamp = &timestamp.Timestamp{Seconds: time.Now().Unix()}

	l.info = append(l.info, info)

	return &empty.Empty{}, nil
}

func (l *Log)Tail(_ *empty.Empty, stream push.LogService_TailServer) error {
	log.Printf("tail:")

	i := len(l.info) - *l.tail
	if i < 0 {
		i = 0
	}

	for _, info := range l.info[i:] {
		if err := stream.Send(info); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	server := push.NewServer()

	_log := NewLog()
	_log.tail = flag.Int("tail", 10, "The 'tail' length")

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterLogServiceServer(s, _log)
	})
}

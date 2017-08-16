package push

import (
	"fmt"
	"google.golang.org/grpc"
)

type Options struct {
	Host *string
	Port *int
	ApiKey *string
}

func (o Options)Dial() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	return grpc.Dial(fmt.Sprintf("%s:%d", *o.Host, *o.Port), opts...)
}

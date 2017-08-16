package main

import (
	"github.com/nokamoto/push-go/proto"
	"flag"
	"google.golang.org/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

type App struct {}

func (a *App)SetFirebaseCloudMessaging(_ context.Context, app *push.FirebaseCloudMessagingApp) (*google_protobuf.Empty, error) {
	return nil, nil
}

func (a *App)GetFirebaseCloudMessaging(_ context.Context, _ *google_protobuf.Empty) (*push.FirebaseCloudMessagingApp, error) {
	return nil, nil
}

func main() {
	server := push.NewServer()

	app := new(App)

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterAppServiceServer(s, app)
	})
}

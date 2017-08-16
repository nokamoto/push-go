package main

import (
	"github.com/nokamoto/push-go/proto"
	"flag"
	"google.golang.org/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"log"
)

type App struct {
	firebase *push.FirebaseCloudMessagingApp
}

func (a *App)SetFirebaseCloudMessaging(_ context.Context, app *push.FirebaseCloudMessagingApp) (*google_protobuf.Empty, error) {
	log.Printf("set: %v", app)
	a.firebase = app
	return new(google_protobuf.Empty), nil
}

func (a *App)GetFirebaseCloudMessaging(_ context.Context, _ *google_protobuf.Empty) (*push.FirebaseCloudMessagingApp, error) {
	log.Printf("get:")
	res := new(push.FirebaseCloudMessagingApp)
	if a.firebase != nil {
		res = a.firebase
	}
	log.Printf("%v", res)
	return res, nil
}

func main() {
	server := push.NewServer()

	app := new(App)

	flag.Parse()

	server.Run(func(s *grpc.Server) {
		push.RegisterAppServiceServer(s, app)
	})
}

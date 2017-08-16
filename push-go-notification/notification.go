package main

import (
	"flag"
	"github.com/nokamoto/push-go/proto"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"errors"
)

type Notification struct {
	app *grpc.ClientConn
	endpoint *grpc.ClientConn
}

func (n *Notification)post(apikey string, payload map[string]interface{}) error {
	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(j))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", apikey))

	client := &http.Client{}

	log.Printf("req: %v", req)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	log.Printf("res: %v", res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Printf("body: %v", string(body))

	// todo: parse the response body and update endpoints

	return nil
}

func (n *Notification)Send(_ context.Context, notification *push.FirebaseCloudMessagingNotification) (*google_protobuf.Empty, error) {
	client := push.NewAppServiceClient(n.app)
	empty := &google_protobuf.Empty{}

	app, err := client.GetFirebaseCloudMessaging(context.Background(), empty)
	if err != nil {
		return nil, err
	}

	var payload map[string]interface {}
	if err = json.Unmarshal([]byte(notification.Payload.Json), &payload); err != nil {
		return nil, err
	}

	tokens := []string{}
	topics := []string{}

	for _, e := range notification.Endpoint {
		if len(e.Token) != 0 {
			tokens = append(tokens, e.Token)
		}
		if len(e.Topic) != 0 {
			topics = append(topics, e.Topic)
		}
	}

	if (len(tokens) > 0 && len(topics) > 0) || (len(topics) > 1) || (len(tokens) > 1000) {
		return nil, errors.New("multiple requests are not available")
	}

	if len(topics) == 1 {
		payload["to"] = topics[0]
		err = n.post(app.ApiKey, payload)
	} else {
		payload["registration_ids"] = tokens
		err = n.post(app.ApiKey, payload)
	}

	return empty, err
}

func main() {
	server := push.NewServer()

	notification := new(Notification)

	app := push.Options{
		Host: flag.String("app.host", "localhost", "The app server host"),
		Port: flag.Int("app.port", 8000, "The app server port"),
		ApiKey: flag.String("app.apikey", "default", "The app request 'x-api-key'"),
	}

	endpoint := push.Options{
		Host: flag.String("endpoint.host", "localhost", "The endpoint server host"),
		Port: flag.Int("endpoint.port", 8000, "The endpoint server port"),
		ApiKey: flag.String("endpoint.apikey", "default", "The endpoint request 'x-api-key'"),
	}

	flag.Parse()

	conn, err := app.Dial()
	if err != nil {
		log.Fatal(err.Error())
	}

	notification.app = conn

	conn, err = endpoint.Dial()
	if err != nil {
		log.Fatal(err.Error())
	}

	notification.endpoint = conn

	server.Run(func(s *grpc.Server) {
		push.RegisterFirebaseCloudMessagingServiceServer(s, notification)
	})
}

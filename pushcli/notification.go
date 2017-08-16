package main

import "github.com/nokamoto/push-go/proto"

type Notification struct {}

func (n Notification)Name() []string {
	return []string{"notification", "n"}
}

func (n Notification)SubCommands() []SubCommand {
	return []SubCommand{
		NewNotificationTo([]string{"topic", "t"}, func(n *push.FirebaseCloudMessagingNotification, topic string){
			n.Topic = topic
		}),
		NewNotificationTo([]string{"condition", "c"}, func(n *push.FirebaseCloudMessagingNotification, condition string){
			n.Condition = condition
		}),
		NewNotificationTo([]string{"endpoint", "e"}, func(n *push.FirebaseCloudMessagingNotification, endpoint string){
			n.Endpoint = []*push.FirebaseCloudMessagingEndpoint{
				{Token: endpoint},
			}
		}),
	}
}


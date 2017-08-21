package main

type Subscription struct {}

func (s Subscription)Name() []string {
	return []string{"subscription", "s"}
}

func (s Subscription)SubCommands() []SubCommand {
	return []SubCommand{SubscriptionSubscribe{}, SubscriptionUnsubscribe{}, SubscriptionGet{}}
}



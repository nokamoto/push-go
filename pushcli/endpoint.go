package main

type Endpoint struct {}

func (e Endpoint)Name() []string {
	return []string{"endpoint", "e"}
}

func (e Endpoint)SubCommands() []SubCommand {
	return []SubCommand{EndpointSet{}, EndpointDelete{}, EndpointGet{}, EndpointUpdate{}}
}

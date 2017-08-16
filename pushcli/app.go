package main

type App struct {}

func (a App)Name() []string {
	return []string{"app", "a"}
}

func (a App)SubCommands() []SubCommand {
	return []SubCommand{AppGet{}, AppSet{}}
}

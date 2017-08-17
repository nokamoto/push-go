package main

type Log struct {}

func (l Log)Name() []string {
	return []string{"log", "l"}
}

func (l Log)SubCommands() []SubCommand {
	return []SubCommand{LogInfo{}, LogTail{}}
}


package main

import (
	"fmt"
	"github.com/nokamoto/push-go/proto"
)

type _MainCommand interface {
	Name() []string

	SubCommands() []SubCommand
}

func run(c _MainCommand, opts push.Options, args []string) error {
	commands := c.SubCommands()

	if len(args) < 1 {
		fmt.Printf("Usage: %v [arguments]\n", c.Name()[0])
		return nil
	}

	for _, command := range commands {
		if matchSubCommandName(command, args[0]) {
			if err := command.Run(opts, args[1:]); err == nil {
				return nil
			} else {
				return err
			}
		}
	}

	fmt.Printf("%v: unknown subcommand \"%v\"\n", c.Name()[0], args[0])

	return nil
}

func matchMainCommandName(c _MainCommand, s string) bool {
	for _, n := range c.Name() {
		if n == s {
			return true
		}
	}
	return false
}

func allCommands() []_MainCommand {
	return []_MainCommand{App{}, Endpoint{}, Notification{}, Log{}, Subscription{}}
}

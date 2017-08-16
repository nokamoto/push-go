package main

import "github.com/nokamoto/push-go/proto"

type SubCommand interface {
	Name() []string

	Run(push.Options, []string) error
}

func matchSubCommandName(c SubCommand, s string) bool {
	for _, n := range c.Name() {
		if n == s {
			return true
		}
	}
	return false
}

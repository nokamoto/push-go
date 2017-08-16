package main

type SubCommand interface {
	Name() []string

	Run(Options, []string) error
}

func matchSubCommandName(c SubCommand, s string) bool {
	for _, n := range c.Name() {
		if n == s {
			return true
		}
	}
	return false
}

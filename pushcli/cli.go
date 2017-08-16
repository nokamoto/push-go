package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	opts := Options{
		Host: flag.String("host", "localhost", "The server host"),
		Port: flag.Int("port", 8000, "The server port"),
		ApiKey: flag.String("apikey", "default", "The request 'x-api-key'"),
	}

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("Usage: %v [flags] command [arguments]\n", os.Args[0])
		os.Exit(0)
	}

	commands := allCommands()

	for _, command := range commands {
		if matchMainCommandName(command, flag.Arg(0)) {
			if err := run(command, opts, flag.Args()[1:]); err == nil {
				os.Exit(0)
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}

	fmt.Printf("%v: unknown subcommand \"%v\"\n", os.Args[0], flag.Args()[0])

	os.Exit(0)
}

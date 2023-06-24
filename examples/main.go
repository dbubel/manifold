package main

import (
	"fmt"
	"github.com/dbubel/manifold/examples/mocks"
	"github.com/mitchellh/cli"
	"os"
)

func main() {
	c := cli.NewCLI("cohesion-content-server", "")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"serve": func() (cli.Command, error) {
			return &mocks.ServeCommand{}, nil
		},
		"consume": func() (cli.Command, error) {
			return &mocks.ConsumeCommand{}, nil
		},
		"produce": func() (cli.Command, error) {
			return &mocks.ProduceCommand{}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		fmt.Println("Error running serve command")
	}

}

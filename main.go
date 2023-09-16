package main

import (
	"github.com/dbubel/manifold/cmd/server"
	"github.com/dbubel/manifold/examples/mocks"
	"github.com/dbubel/manifold/pkg/logging"
	"github.com/mitchellh/cli"
	"os"
)

func main() {
	l := logging.New(logging.DEBUG)

	c := cli.NewCLI("manifold", "")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"serve": func() (cli.Command, error) {
			return &server.ManifoldServerCmd{}, nil
		},
		"consumer": func() (cli.Command, error) {
			return &mocks.ConsumeCommand{}, nil
		},
		"producer": func() (cli.Command, error) {
			return &mocks.ProduceCommand{
				os.Args[1:],
			}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		l.WithFields(map[string]interface{}{"error": err.Error()}).Error("error running serve command")
	}
}

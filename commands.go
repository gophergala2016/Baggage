package main

import (
	"github.com/gophergala2016/Baggage/command"
	"github.com/mitchellh/cli"
	"os"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Ui: ui,
			}, nil
		},
		"new": func() (cli.Command, error) {
			return &command.NewCommand{
				Ui: ui,
			}, nil
		},
	}

}

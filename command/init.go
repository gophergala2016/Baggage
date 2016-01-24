package command

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"strings"
)

type InitCommand struct {
	Ui cli.Ui
}

func (c *InitCommand) Help() string {
	helpText := `
	Usage: Baggage init
	`

	return strings.TrimSpace(helpText)
}

func (c *InitCommand) Run(args []string) int {
	if _, err := os.Stat("./baggage.dat"); os.IsExist(err) {
		fmt.Println("true:baggage.dat")
		os.Remove("./baggage.dat")
	} else {
		fmt.Println("false:baggage.dat")
	}

	return 0
}

func (c *InitCommand) Synopsis() string {
	return "Bagage init"
}

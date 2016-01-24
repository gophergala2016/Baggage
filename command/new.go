package command

import (
	_ "flag"
	"fmt"
	"github.com/gophergala2016/Baggage/parser"
	"github.com/mitchellh/cli"
	"strings"
)

/*type Project struct {

}*/

type NewCommand struct {
	Ui cli.Ui
}

func (c *NewCommand) Help() string {
	helpText := `
	Usage: Baggage new name
	`
	return strings.TrimSpace(helpText)
}

func (c *NewCommand) Run(args []string) int {
	c.Ui.Output(fmt.Sprintln("New Command"))
	for i, v := range args {
		fmt.Println(i, ", ", v)
	}
	prj := parser.NewParser("./baggage.xml")
	prj.Parse()
	return 0
}

func (c *NewCommand) Synopsis() string {
	return "NewCommand::Synopsis"
}

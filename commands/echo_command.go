package commands

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type EchoCommand struct {
	name string
}

func (command EchoCommand) GetName() string {
	return "echo"
}

func (command EchoCommand) String() string {
	return "<Command 'echo'>"
}

func (command EchoCommand) Execute(args []string) {
	fmt.Fprintln(color.Output, strings.Join(args, " "))
}

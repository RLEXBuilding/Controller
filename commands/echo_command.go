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

func (EchoCommand) IsWIP() bool {
	return false
}
func (EchoCommand) IsIllegal() bool {
	return false
}
func (EchoCommand) RequiresSU() bool {
	return false
}
func (command EchoCommand) GetDescription() string {
	return "Prints a string"
}

func (command EchoCommand) String() string {
	return "<Command 'echo'>"
}

func (command EchoCommand) Execute(kill chan bool, args []string) {
	fmt.Fprintln(color.Output, strings.Join(args, " "))
}

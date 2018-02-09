package commands

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// EchoCommand is a command, which prints a message to the output
type EchoCommand struct {
	name string
}

// GetName is the function, which returns the name of the command
func (command EchoCommand) GetName() string {
	return "echo"
}

// GetDescription is a command, which returns the description
func (command EchoCommand) GetDescription() string {
	return "Prints a string"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (EchoCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (EchoCommand) RequiresSU() bool {
	return false
}

func (command EchoCommand) String() string {
	return "<Command 'echo'>"
}

// Execute is a function which executes the command
func (command EchoCommand) Execute(kill chan bool, args []string) {
	fmt.Fprintln(color.Output, strings.Join(args, " "))
}

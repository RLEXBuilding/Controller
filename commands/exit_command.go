package commands

import (
	"fmt"
	"os"
)

// ExitCommand is a command, which says bye bye"
type ExitCommand struct {
	name string
}

// GetName is the function, which returns the name of the command
func (command ExitCommand) GetName() string {
	return "exit"
}

// GetDescription is the function, which returns the description of the command
func (command ExitCommand) GetDescription() string {
	return "Says Bye Bye"
}

func (command ExitCommand) String() string {
	return "<Command 'exit'>"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (ExitCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (ExitCommand) RequiresSU() bool {
	return false
}

// Execute is the function, which executes the command
func (command ExitCommand) Execute(kill chan bool, args []string) {
	fmt.Println("See You later Alligator")
	os.Exit(0)
}

package commands

import (
	"fmt"
	"os"
)

type ExitCommand struct {
	name string
}

func (command ExitCommand) GetName() string {
	return "exit"
}

func (command ExitCommand) GetDescription() string {
	return "Says Bye Bye"
}

func (command ExitCommand) String() string {
	return "<Command 'exit'>"
}
func (ExitCommand) IsWIP() bool {
	return false
}
func (ExitCommand) IsIllegal() bool {
	return false
}
func (ExitCommand) RequiresSU() bool {
	return false
}
func (command ExitCommand) Execute(kill chan bool, args []string) {
	fmt.Println("See You later Alligator")
	os.Exit(0)
}

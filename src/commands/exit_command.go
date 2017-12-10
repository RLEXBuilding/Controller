package commands

import "fmt"
import "os"

type  ExitCommand struct {
	name string
}

func (command ExitCommand) GetName() string {
	return "exit"
}

func (command ExitCommand) String() string {
	return "<Command 'exit'>"
}

func (command ExitCommand) Execute(args []string) {
	fmt.Println("See You later Alligator")
	os.Exit(0)
}

package commands

import (
	"fmt"
	"os"
	"time"
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

func (command ExitCommand) Execute(kill chan bool, args []string) {
	fmt.Println("See You later Alligator")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

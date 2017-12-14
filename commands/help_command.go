package commands

import (
	"fmt"

	"github.com/fatih/color"
)

type HelpCommand struct {
	name string
}

func (command HelpCommand) GetName() string {
	return "help"
}

func (command HelpCommand) GetDescription() string {
	return "Shows a help ui"
}

func (command HelpCommand) String() string {
	return "<Command 'help'>"
}

func (command HelpCommand) Execute(kill chan bool, args []string) {
	fmt.Println("--- Help ---")
	for _, element := range commands {
		fmt.Fprintln(color.Output, element.GetName()+" | "+element.GetDescription())
	}
}

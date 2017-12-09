package commands

import "fmt"

type HelpCommand struct {
	name string
}

func (command HelpCommand) GetName() string {
	return "help"
}

func (command HelpCommand) String() string {
	return "<Command 'help'>"
}

func (command HelpCommand) Execute(args []string) {
	fmt.Println("HELP")
	for _, element := range commands {
		fmt.Println(element.GetName())
	}
}

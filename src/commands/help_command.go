package commands

import (
	"fmt"
	"strconv"
)

type HelpCommand struct {
	name string
}

func (command HelpCommand) GetName() string {
	return "help"
}

func (command HelpCommand) Execute(args []string) {
	fmt.Println("HELP")
	for i, element := range commands {
		fmt.Println(strconv.Itoa(i) + ":" + element.GetName())
	}
}

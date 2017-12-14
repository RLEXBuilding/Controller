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
	explanation := "| " + color.YellowString("\u2588"+" = Could be illegal")
	explanation += color.RedString("\u2588"+" = No Permission") + " | "
	explanation += color.CyanString("\u2588"+" = Currently not working") + " | "

	fmt.Println(explanation)
	fmt.Println("------------")
	for _, element := range commands {
		fmt.Fprintln(color.Output, element.GetName()+" | "+element.GetDescription())
	}
}

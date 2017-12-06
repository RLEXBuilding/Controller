package commands

import "strings"
import "fmt"

type Command interface {
	GetName() string
	Execute(args []string)
}

var commands []Command

func InitializeCommands() {
	commands = append(commands, new(HelpCommand))
}

func RunCommand(cmd string) {
	args := strings.Fields(cmd)
	var element = getElementByString(cmd)
	if element == nil {
		fmt.Println("Cannot resolve this command. Type help for a help gui")
	} else {
		element.Execute(append(args[:0], args[0+1:]...))
	}
}

func getElementByString(cmd string) Command {
	for _, element := range commands {
		if strings.EqualFold(element.GetName(), cmd) {
			return element
		}
	}
	return nil
}

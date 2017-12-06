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
	args := strings.Split(cmd, " ")
	var element = getElementByString(args[0])
	if element == nil {
		fmt.Println("Cannot resolve this command. Type help for a help gui")
	} else {
		element.Execute(append(args[:0], args[0+1:]...))
	}
}

func getElementByString(cmd string) Command {
	fmt.Println(cmd)
	for _, element := range commands {
		fmt.Println(element.GetName())
		if strings.EqualFold(element.GetName(), cmd) {
			return element
		}
	}
	return nil
}

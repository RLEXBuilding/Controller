package commands

import (
	"fmt"
	"strings"
)

type Command interface {
	GetName() string
	Execute(args []string)
}

var commands []Command

func InitializeCommands() {
	commands = append(commands, new(ExitCommand))
	commands = append(commands, new(HelpCommand))
	commands = append(commands, new(ShellCommand))
	commands = append(commands, new(WhoisCommand))
	commands = append(commands, new(PortScanCommand))
	commands = append(commands, new(SshBruteForceCommand))
}

func RunCommand(cmd string) {
	args := strings.Split(strings.TrimSpace(cmd), " ")
	element := getElementByString(args[0])
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

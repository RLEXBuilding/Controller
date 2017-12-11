package commands

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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
	commands = append(commands, new(ListFilesCommand))
	commands = append(commands, new(HackImitateCommand))
}

func RunCommand(cmd string) {
	args := strings.Split(strings.TrimSpace(cmd), " ")
	element := getElementByString(args[0])
	if element == nil {
		fmt.Fprintln(color.Output, "Cannot resolve this command. Type "+color.YellowString("help")+" for a help gui")
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

func getStringRepeatedByInt(str string, repeated int) string {
	str2 := ""
	i := 1
	for i <= repeated {
		str2 += str
		i++
	}

	return str2
}

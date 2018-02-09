package commands

import (
	"fmt"
	"strings"

	"os"
	"os/signal"

	"github.com/fatih/color"
)

// Command is a type, which defines a command
type Command interface {
	GetName() string
	GetDescription() string
	Execute(kill chan bool, args []string)
	IsWIP() bool
	RequiresSU() bool
}

var commands []Command

// InitializeCommands is a function, which registers all commands
func InitializeCommands() {
	commands = append(commands, new(ExitCommand))
	commands = append(commands, new(HelpCommand))
	commands = append(commands, new(EchoCommand))
	commands = append(commands, new(ShellCommand))
	commands = append(commands, new(WgetCommand))
	commands = append(commands, new(WhoisCommand))
	commands = append(commands, new(PortScanCommand))
	commands = append(commands, new(SSHBruteForceCommand))
	commands = append(commands, new(ListFilesCommand))
	commands = append(commands, new(CheckPasswordCommand))
	commands = append(commands, new(HackImitateCommand))
}

// RunCommand is a function, which runs a command by arguments
func RunCommand(args []string) {
	element := getElementByString(args[0])
	if element == nil {
		fmt.Fprintln(color.Output, "Cannot resolve this command. Type "+color.YellowString("help")+" for a help gui")
	} else {
		RunCommandAsync(element, args)
	}
}

// RunCommandAsync is a function, which runs a command without sync
func RunCommandAsync(command Command, args []string) {
	kill := make(chan bool, 1)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		for range signals {
			kill <- true
		}
	}()
	command.Execute(kill, args[1:])
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

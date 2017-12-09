package commands

import "strings"
import "fmt"

type Command interface {
	GetName() string
	Execute(args []string)
}

var commands []Command

func InitializeCommands() {
	commands = append(commands, new(ExitCommand))
	commands = append(commands, new(HelpCommand))
}

func RunCommand(cmd string) {
	args := strings.Split(strings.TrimSpace(cmd), " ")
	//fmt.Printf("Input: %s\n", args[0])
	element := getElementByString(args[0])
	//fmt.Printf("Result: %s\n", element)
	if element == nil {
		fmt.Println("Cannot resolve this command. Type help for a help gui")
	} else {
		element.Execute(append(args[:0], args[0+1:]...))
	}
}

func getElementByString(cmd string) Command {
	for _, element := range commands {
//		fmt.Printf("           Input: '%s'\n", cmd)
//		fmt.Printf("Querying Command: %s\n",element)
//		fmt.Printf("            Name: '%s'\n", element.GetName())
		if strings.EqualFold(element.GetName(), cmd) {
			return element
		}
	}
	return nil
}

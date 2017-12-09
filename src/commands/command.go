package commands

import (
	"fmt"
	"strconv"
	"strings"
)

type Command interface {
	GetName() string
	Execute(args []string)
}

var commands []Command

func InitializeCommands() {
	commands = append(commands, new(HelpCommand))
	commands = append(commands, new(WhoisCommand))
}

func RunCommand(cmd string) {

	args := deleteEmptyEntrysInSlice(strings.Split(""+cmd, " "))

	for i, arg := range args {
		fmt.Println(strconv.Itoa(i) + "-" + arg)
	}

	var element = getElementByString(args[0])
	if element == nil {
		fmt.Println("Cannot resolve this command. Type help for a help gui")
	} else {
		element.Execute(append(args[:0], args[0+1:]...))
	}
}

func getElementByString(cmd string) Command {
	fmt.Println(cmd)
	for i, element := range commands {
		fmt.Println(strconv.Itoa(i) + ":" + element.GetName())
		if strings.EqualFold(element.GetName(), cmd) {
			fmt.Println("Resolved " + element.GetName())
			return element
		}
	}
	return nil
}

func deleteEmptyEntrysInSlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
			fmt.Println(str + "-YES")
		} else {
			fmt.Println(str + "-NOPE")
		}
	}
	return r
}

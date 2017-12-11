package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

type ShellCommand struct {
	name string
}

func (command ShellCommand) GetName() string {
	return "shell"
}

func (command ShellCommand) String() string {
	return "<Command 'shell'>"
}

func (command ShellCommand) Execute(args []string) {
	if len(args) < 1 {
		fmt.Println("Missing arguments.")
		return
	}
	cmd := strings.Join(args[1:], " ")
	res, _ := exec.Command(args[0], cmd).Output()
	fmt.Println("= " + string(res))
}

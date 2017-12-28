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
func (ShellCommand) IsWIP() bool {
	return false
}
func (ShellCommand) RequiresSU() bool {
	return false
}

func (command ShellCommand) GetDescription() string {
	return "Executes a shell command"
}

func (command ShellCommand) String() string {
	return "<Command 'shell'>"
}

func (command ShellCommand) Execute(kill chan bool, args []string) {
	if len(args) < 1 {
		fmt.Println("Missing arguments.")
		return
	}
	cmd := strings.Join(args[1:], " ")
	res, _ := exec.Command(args[0], cmd).Output()
	fmt.Println("= " + string(res))
}

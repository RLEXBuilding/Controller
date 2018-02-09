package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

// ShellCommand is the command which executes an shell command in the default shell of the operation system
type ShellCommand struct {
	name string
}

// GetName is the function which returns the name of the shell command
func (command ShellCommand) GetName() string {
	return "shell"
}

// GetDescription returns the description of the shell command
func (command ShellCommand) GetDescription() string {
	return "Executes a shell command"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (ShellCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (ShellCommand) RequiresSU() bool {
	return false
}

func (command ShellCommand) String() string {
	return "<Command 'shell'>"
}

// Execute executes the shell command
func (command ShellCommand) Execute(kill chan bool, args []string) {
	if len(args) < 1 {
		fmt.Println("Missing arguments.")
		return
	}
	cmd := strings.Join(args[1:], " ")
	res, _ := exec.Command(args[0], cmd).Output()
	fmt.Println("= " + string(res))
}

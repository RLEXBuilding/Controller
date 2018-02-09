package commands

import (
	"fmt"

	"io/ioutil"
	"os"
	"os/user"
	"runtime"

	"github.com/fatih/color"
)

// HelpCommand is a command which shows a help ui
type HelpCommand struct {
	name string
}

// GetName returns the name of the command
func (command HelpCommand) GetName() string {
	return "help"
}

// GetDescription returns the description of the command
func (command HelpCommand) GetDescription() string {
	return "Shows a help ui"
}

func (command HelpCommand) String() string {
	return "<Command 'help'>"
}

func checkSU() bool {
	user, err := user.Current()
	if err != nil {
		return false
	}
	if runtime.GOOS == "linux" && user.Uid == "0" && user.Gid == "0" {
		return true
	}
	if runtime.GOOS == "windows" {
		_, err := ioutil.ReadFile(os.Getenv("SYSTEMROOT") + "\\" + "Windows\\System32\\user32.dll")
		return err != nil
	}
	return false
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (HelpCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (HelpCommand) RequiresSU() bool {
	return false
}

// Execute is the function which executes the command
func (command HelpCommand) Execute(kill chan bool, args []string) {
	fmt.Println("+---     Help      ---+")
	fmt.Fprintln(color.Output, color.RedString("\u2588"+" = No Permission"))
	fmt.Fprintln(color.Output, color.CyanString("\u2588"+" = WIP could be dysfunctional"))
	fmt.Println("+--- Key Shortcuts ---+")
	fmt.Println("^C | Abort current command")
	fmt.Println("+---   Commands    ---+")

	for _, element := range commands {
		text := element.GetName() + " | " + element.GetDescription()
		if element.IsWIP() {
			fmt.Fprintln(color.Output, color.CyanString(text))
		} else if element.RequiresSU() && !checkSU() {
			fmt.Fprintln(color.Output, color.RedString(text))
		} else {
			fmt.Println(element.GetName() + " | " + element.GetDescription())
		}
	}
}

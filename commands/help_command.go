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

func (command HelpCommand) GetName() string {
	return "help"
}

func (command HelpCommand) GetDescription() string {
	return "Shows a help ui"
}

func (command HelpCommand) String() string {
	return "<Command 'help'>"
}

func buildBadges(cmd Command) string {
	badges := ""
	if cmd.IsWIP() {
		badges += color.CyanString("[WIP]")
	}
	if cmd.IsIllegal() {
		badges += color.YellowString("[ILLEGAL]")
	}
	if cmd.RequiresSU() {
		badges += color.RedString("[SU]")
	}
	return badges
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

func (HelpCommand) IsWIP() bool {
	return false
}
func (HelpCommand) IsIllegal() bool {
	return false
}
func (HelpCommand) RequiresSU() bool {
	return false
}
func (command HelpCommand) Execute(kill chan bool, args []string) {
	fmt.Println("+---     Help      ---+")
	fmt.Fprintln(color.Output, color.YellowString("\u2588"+" = Could be illegal"))
	fmt.Fprintln(color.Output, color.RedString("\u2588"+" = No Permission"))
	fmt.Fprintln(color.Output, color.CyanString("\u2588"+" = WIP could be dysfunctional"))
	fmt.Println("+--- Key Shortcuts ---+")
	fmt.Println("^C | Abort current command")
	fmt.Println("+---   Commands    ---+")

	for _, element := range commands {
		text := element.GetName() + buildBadges(element) + "|" + element.GetDescription()
		if element.IsWIP() {
			fmt.Fprintln(color.Output, color.CyanString(text))
		} else if element.IsIllegal() {
			fmt.Fprintln(color.Output, color.YellowString(text))
		} else if element.RequiresSU() && !checkSU() {
			fmt.Fprintln(color.Output, color.RedString(text))
		} else {
			fmt.Println(element.GetName() + " | " + element.GetDescription())
		}
	}
}

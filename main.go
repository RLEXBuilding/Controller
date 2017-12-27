package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/RLEXBuilding/Controller/commands"
	"github.com/RLEXBuilding/Controller/util"
	"github.com/fatih/color"
)

func main() {
	commands.InitializeCommands()
	fmt.Fprintln(color.Output, "Welcome on the Controller of "+color.MagentaString("magic"))
	fmt.Println("We don't support illegal actions. It's your choice :).")
	fmt.Println()
	fmt.Println("Operation System: " + runtime.GOOS)
	fmt.Println("Architecture: " + runtime.GOARCH)
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for true {
	OUTER:
		fmt.Fprint(color.Output, color.GreenString("> "))
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		arr, finished, err := util.ParseQuotes(text)
		if err != nil {
			fmt.Println(err)
			goto OUTER
		}
		for !finished {
			nl, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			text = text + nl
			arr, finished, err = util.ParseQuotes(text)
		}
		if len(arr) == 0 {
			goto OUTER
		}
		commands.RunCommand(arr)
	}
}

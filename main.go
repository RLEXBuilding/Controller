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
			fmt.Fprintln(color.Output, err)
			continue
		}
		arr, finished, err := util.ParseQuotes(text)
		if err != nil {
			fmt.Print(err)
			continue
		}
		for !finished {
			nl, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(color.Output, err)
				goto OUTER
			}
			text = text + nl
			arr, finished, err = util.ParseQuotes(text)
		}
		if len(arr) == 0 {
			continue
		}
		commands.RunCommand(arr)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"./commands"
	"github.com/fatih/color"
)

func main() {
	commands.InitializeCommands()
	fmt.Fprintln(color.Output, "Welcome on the Controller of "+color.MagentaString("magic"))
	fmt.Println("We don't support illegal actions. It's your choice :).")
	fmt.Println()
	fmt.Println("OS: " + runtime.GOOS)
	fmt.Println("ARCH: " + runtime.GOARCH)
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Fprint(color.Output, color.GreenString("> "))
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		commands.RunCommand(text)
	}
}

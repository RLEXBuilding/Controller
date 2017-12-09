package main

import (
	"bufio"
	"fmt"
	"os"

	commands "./commands"
)

func main() {
	commands.InitializeCommands()
	fmt.Println("Welcome on the Controller of magic")
	fmt.Println("We don't support illegal actions. It's your choice :).")
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		commands.RunCommand(text)
	}
}

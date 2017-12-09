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
	fmt.Println()
	var count int
	fmt.Scan(&count)
	reader := bufio.NewReader(os.Stdin)
	//m := make(map[string]string)
	for true {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("$ " + text)
		commands.RunCommand(text)
		//value := strings.Fields(text)
		//m[value[0]] = value[1]
	}
}

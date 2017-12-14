package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anvie/port-scanner"
	"github.com/fatih/color"
)

type PortScanCommand struct {
	name string
}

func (command PortScanCommand) GetName() string {
	return "portscan"
}

func (command PortScanCommand) GetDescription() string {
	return color.YellowString("Scans a port range")
}

func (command PortScanCommand) String() string {
	return "<Command 'portscan'>"
}

func (command PortScanCommand) Execute(kill chan bool, args []string) {
	/*
		This command is not finished. If you want to help:
		- Please add a system with arguments like "-tag"
		  > "-async" argument
		  > "-listClosed" argument
		  > "-dontListOpened" argument
		  > "-onlyList=open,closed" argument(or something like this)
		  > "-asList" argument(should be displayed: 80,81)
	*/

	if len(args) < 3 {
		fmt.Println("portscan <address> <port-from> <port-to>")
		return
	}

	address := args[0]
	from, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("The second argument is not a valid number")
		return
	}
	to, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("The third argument is not a valid number")
		return
	}

	ps := portscanner.NewPortScanner(address, 2*time.Second, 5)

	fmt.Printf("scanning port %d-%d...\n", from, to)

	for i := from; i < to; i++ {
		if ps.IsOpen(i) {
			fmt.Fprint(color.Output, " ", i, color.GreenString(" [open]"))
			fmt.Println("  -  ", ps.DescribePort(i))
		}
	}
}

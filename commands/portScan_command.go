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
	if len(args) < 1 {
		fmt.Println("portscan <address> [port-from-inclusive] [port-to-exclusive] [timeout-in-milliseconds]")
		return
	}

	address := args[0]
	from := 0
	to := 65535
	var timeout = 2000 * time.Millisecond
	if len(args) >= 3 {
		var err error
		from, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("The second argument is not a valid number")
			return
		}
		to, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("The third argument is not a valid number")
			return
		}
	}
	if len(args) >= 4 {
		iTimeout, err := strconv.Atoi(args[3])
		if err != nil || to <= 0 {
			fmt.Println("Invalid timeout.")
			return
		}
		timeout = time.Duration(iTimeout) * time.Millisecond
	}

	ps := portscanner.NewPortScanner(address, timeout, 5)

	// get opened port
	fmt.Printf("scanning port %d-%d...\n", from, to)

	for port := from; port < to; port++ {
		select {
		case <-kill:
			fmt.Println("\rAborted.")
			return

		default:
			fmt.Printf("\rCurrent port: %5d", port)
			if ps.IsOpen(port) {
				fmt.Printf("\r %d [open]  -  %s\n", port, ps.DescribePort(port))
			}
		}
	}
}

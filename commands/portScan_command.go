package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anvie/port-scanner"
)

type PortScanCommand struct {
	name string
}

func (command PortScanCommand) GetName() string {
	return "portscan"
}

func (command PortScanCommand) String() string {
	return "<Command 'portscan'>"
}

func (command PortScanCommand) Execute(args []string) {
	if len(args) < 3 {
		fmt.Println("portscan <address> <port-from> <port-to>")
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

	// get opened port
	fmt.Printf("scanning port %d-%d...\n", from, to)

	openedPorts := ps.GetOpenedPort(from, to)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -  ", ps.DescribePort(port))
	}
}

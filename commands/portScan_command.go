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
	from := 1
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
			progressBarMax := 10.0
			progressBarStatus := ((float64(port) - float64(from)) / (float64(to) - float64(from))) * progressBarMax
			progressBarPercent := ((float64(port) - float64(from)) / (float64(to) - float64(from))) * 100.0

			progressBar := "["
			for i := 0; i <= int(progressBarStatus); i++ {
				progressBar += color.GreenString("\u2588")
			}

			remaining := progressBarMax - progressBarStatus

			for i2 := 0; i2 <= int(remaining); i2++ {
				progressBar += color.HiWhiteString("\u2588")
			}
			progressBar += "] " + strconv.FormatFloat(progressBarPercent, 'f', 1, 64)
			fmt.Fprintf(color.Output, "\rCurrent port: %5d %s", port, progressBar+string('\u0025'))
			if ps.IsOpen(port) {
				fmt.Printf("\r %d [open]  -  %s\t\t\t\t\n", port, ps.DescribePort(port))
			}
		}
	}
}

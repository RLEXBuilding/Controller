package commands

import (
	"fmt"
	"strings"

	whois "github.com/undiabler/golang-whois"
)

// WhoisCommand is a command, which gives you some data about domains
type WhoisCommand struct {
	name string
}

// GetName is the function which returns the name of the command
func (command WhoisCommand) GetName() string {
	return "whois"
}

// GetDescription returns the description of the command
func (command WhoisCommand) GetDescription() string {
	return "Domain Whois"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (WhoisCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (WhoisCommand) RequiresSU() bool {
	return false
}

func (command WhoisCommand) String() string {
	return "<Command 'whois'>"
}

// Execute is a function which executes the command
func (command WhoisCommand) Execute(kill chan bool, args []string) {
	/*
		This command is unfinished.
		Please add sysadmin etc. to this domain whois.
		And please add a ip whois
	*/
	if len(args) != 2 {
		fmt.Println("Usage: whois domain|ip <address>")
		return
	}
	if strings.EqualFold(args[0], "domain") {
		result, err := whois.GetWhois(args[1])

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)
	} else if strings.EqualFold(args[0], "ip") {
		fmt.Println("Not implemented")
	} else {
		fmt.Println("Usage: domain|ip <address>")
	}

}

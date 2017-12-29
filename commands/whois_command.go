package commands

import (
	"fmt"
	"strings"

	whois "github.com/undiabler/golang-whois"
)

type WhoisCommand struct {
	name string
}

func (WhoisCommand) IsWIP() bool {
	return false
}
func (WhoisCommand) RequiresSU() bool {
	return false
}

func (command WhoisCommand) GetName() string {
	return "whois"
}

func (command WhoisCommand) GetDescription() string {
	return "Domain Whois"
}

func (command WhoisCommand) String() string {
	return "<Command 'whois'>"
}

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

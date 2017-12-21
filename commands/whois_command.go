package commands

import (
	"fmt"

	"github.com/likexian/whois-go"
)

type WhoisCommand struct {
	name string
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
		fmt.Println("whois domain <address>")
		return
	}

	result, err := whois.Whois(args[1])
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

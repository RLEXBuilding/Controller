package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type ListFilesCommand struct {
	name string
}

func (command ListFilesCommand) GetName() string {
	return "listFiles"
}

func (command ListFilesCommand) String() string {
	return "<Command 'listFiles'>"
}

func (command ListFilesCommand) Execute(args []string) {
	dir := ""
	if len(args) > 0 {
		dir = strings.Join(args, " ")
	} else {
		dir = os.Getenv("SYSTEMDRIVE") + "\\"
	}

	explanation := "| " + color.YellowString("\u2588"+" = Directory") + " | "
	explanation += color.RedString("\u2588"+" = Directory without access") + " | "
	explanation += color.CyanString("\u2588"+" = File") + " | "

	fmt.Fprintln(color.Output, explanation)
	listFiles(dir, 0)
}

func listFiles(dir string, spaces int) {

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
			if !strings.EqualFold(file.Name(), "$Recycle.Bin") {
				arr, err := ioutil.ReadDir(filepath.Join(dir, file.Name()))
				if err != nil {
					fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.RedString(file.Name())+" - "+color.RedString(err.Error()))
					err = nil
				} else {
					if len(arr) > 0 {
						fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.YellowString(file.Name()+" {"))
						listFiles(filepath.Join(dir, file.Name()), spaces+2)
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
			fmt.Fprintln(color.Output, getStringRepeatedByInt(" ", spaces)+color.CyanString(file.Name()))
		}
	}
}

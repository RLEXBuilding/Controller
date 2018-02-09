package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/DeepRobin/Controller/util"
	"github.com/fatih/color"
)

// ListFilesCommand is a command which lists files
type ListFilesCommand struct {
	name string
}

// GetName is function which returns the name of the command
func (command ListFilesCommand) GetName() string {
	return "listFiles"
}

// GetDescription returns the description of the command
func (command ListFilesCommand) GetDescription() string {
	return "Lists files"
}

// IsWIP is a function which returns the state which defines if the command is unfinished developed
func (ListFilesCommand) IsWIP() bool {
	return false
}

// RequiresSU is the function which returns if the command needs administrator access
func (ListFilesCommand) RequiresSU() bool {
	return false
}

func (command ListFilesCommand) String() string {
	return "<Command 'listFiles'>"
}

// Execute is a function which executes the command
func (command ListFilesCommand) Execute(kill chan bool, args []string) {
	dir := ""
	arguments := ""
	if len(args) > 0 {
		startIndex, endIndex := util.DetectString(args)
		dir = strings.Join(args, " ")

		if startIndex != -1 || endIndex != -1 {
			dir = string([]rune(dir)[startIndex+1 : endIndex])
		} else {
			dir = strings.Split(strings.Join(args, " "), "-")[0]
			argumentsSplit := strings.Split(strings.Join(args, " "), "-")
			argumentsSplit = argumentsSplit[1:len(argumentsSplit)]
			arguments = "-" + strings.Join(argumentsSplit, " ")
		}
	} else {
		dir = os.Getenv("SYSTEMDRIVE") + "\\"
	}

	explanation := "| " + color.YellowString("\u2588"+" = Directory") + " | "
	explanation += color.RedString("\u2588"+" = Directory without access") + " | "
	explanation += color.CyanString("\u2588"+" = File") + " | "
	dirs := true
	files := true

	fmt.Fprintln(color.Output, explanation)
	if strings.Contains(strings.ToLower(arguments), strings.ToLower("-noDir")) {
		dirs = false
	}

	if strings.Contains(strings.ToLower(arguments), strings.ToLower("-noFile")) {
		files = false
	}

	if strings.Contains(strings.ToLower(arguments), strings.ToLower("-all")) {
		listAllFiles(dir, 0)
	} else {
		listFiles(dir, dirs, files)
	}
}

func listFiles(dir string, listDirs bool, listFiles bool) {

	if !listDirs && !listFiles {
		fmt.Fprintln(color.Output, color.HiMagentaString("Nothing - your filters are not allowing files and directorys"))
		return
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("error - " + err.Error())
	}
	for _, file := range files {
		str := ""
		if file.IsDir() && listDirs {
			_, err := ioutil.ReadDir(filepath.Join(dir, file.Name()))
			if err != nil {
				str += color.RedString(file.Name())
			} else {
				str += color.YellowString(file.Name())
			}
		} else if listFiles {
			str += color.CyanString(file.Name())
		}
		fmt.Fprintln(color.Output, str)
	}
}

func listAllFiles(dir string, spaces int) {

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
						listAllFiles(filepath.Join(dir, file.Name()), spaces+2)
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

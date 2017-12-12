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
		dir = strings.Join(args, " ") // TODO: Add here please a support for the -noDir tags etc. Thanks!
	} else {
		dir = os.Getenv("SYSTEMDRIVE") + "\\"
	}

	explanation := "| " + color.YellowString("\u2588"+" = Directory") + " | "
	explanation += color.RedString("\u2588"+" = Directory without access") + " | "
	explanation += color.CyanString("\u2588"+" = File") + " | "
	dirs := true
	files := true

	fmt.Fprintln(color.Output, explanation)
	if strings.Contains(strings.ToLower(strings.Join(args, " ")), strings.ToLower("-noDir")) {
		dirs = false
	}

	if strings.Contains(strings.ToLower(strings.Join(args, " ")), strings.ToLower("-noFile")) {
		files = false
	}

	if strings.Contains(strings.ToLower(strings.Join(args, " ")), strings.ToLower("-all")) {
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

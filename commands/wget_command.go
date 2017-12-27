package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/RLEXBuilding/Controller/util"
	"github.com/fatih/color"
)

type WgetCommand struct {
	name string
}

func (command WgetCommand) GetName() string {
	return "wget"
}

func (WgetCommand) IsWIP() bool {
	return true
}
func (WgetCommand) IsIllegal() bool {
	return false
}
func (WgetCommand) RequiresSU() bool {
	return false
}
func (command WgetCommand) GetDescription() string {
	return "Downloads a file"
}

func (command WgetCommand) String() string {
	return "<Command 'wget'>"
}

func (command WgetCommand) Execute(kill chan bool, args []string) {

	/*
	  The string detection with multiple strings don't work.
	  Please fix this here
	*/

	fmt.Println("This command is disabled, because buggy")
	return

	url := ""
	path := ""

	urlStartIndex, urlEndIndex := util.DetectString(args)
	url = strings.Join(args, " ")

	if urlStartIndex != -1 || urlEndIndex != -1 {
		url = string([]rune(url)[urlStartIndex+1 : urlEndIndex])
	} else {
		fmt.Println("Please give me a string in the char \"")
		fmt.Println("Correct Syntax: " + "wget <url:string> <path:string>")
	}

	pathArgs := strings.Split(strings.Replace(strings.Join(args, " "), url, "", 1), " ")
	pathStartIndex, pathEndIndex := util.DetectString(pathArgs)
	path = strings.Join(pathArgs, " ")

	if pathStartIndex != -1 || pathEndIndex != -1 {
		path = string([]rune(url)[pathStartIndex+1 : pathEndIndex])
	} else {
		fmt.Println("Please give me a string in the char \"")
	}

	fmt.Fprintln(color.Output, "Started downloading from "+color.CyanString(url)+" to "+color.CyanString(path))
	downloadFile(path, url)

}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

package commands

import (
	"fmt"
	"io"
	"net/http"
	netUrl "net/url"
	"os"
	"strconv"
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

	stringArr, finished, error := util.ParseQuotes(strings.Join(args, " "))
	if error != nil {
		fmt.Println("String parsing error: " + error.Error())
		return
	}
	if !finished {
		fmt.Println("String parsing not finished")
		return
	}
	if len(stringArr) < 1 {
		fmt.Println("We need 2 strings. You have " + strconv.Itoa(len(stringArr)))
		return
	}
	url := stringArr[0]
	u, err := netUrl.Parse(url)
	if err != nil {
		fmt.Println("Error while url parsing: " + err.Error())
		return
	}
	localPath := strings.Replace(u.Path, "/", "", 0)
	if strings.TrimSpace(localPath) == "" {
		localPath = u.Host
	}
	if len(stringArr) == 2 {
		localPath = stringArr[1]
	}

	fmt.Fprintln(color.Output, "Started downloading from "+color.CyanString(url)+" to "+color.CyanString(localPath))
	downloadFile(localPath, url)

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

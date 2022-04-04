//usr/bin/go run $0$@ ; exit

package main

import (
	"fmt"
	"github.com/jkomyno/nanoid"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func checkKey() (exists bool, value string) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	id, err := ioutil.ReadFile(userHomeDir + "/c2c.dat")
	if err != nil {
		fmt.Println("No key found")
		fmt.Println("Creating new key...")
	}
	if len(id) == 0 {
		return false, ""
	} else {
		return true, string(id)
	}
}

func createKey() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	id, err := nanoid.Nanoid(8)
	if err != nil {
		fmt.Println(err)
	}
	writeErr := ioutil.WriteFile(userHomeDir+"/c2c.dat", []byte(id), 0777)
	if err != nil {
		fmt.Println(writeErr)
	}
	fmt.Printf("Created key: %s\n", id)
	fmt.Printf("Web interface: https://cli2cloud.herokuapp.com/#/cli?id=%s\n", id)
}

func sendData(data string, value string) {
	post := url.Values{
		"data": {data},
		"room": {value},
	}
	_, err := http.PostForm("https://cli2cloud.herokuapp.com/api/v1/socket", post)
	// _, err := http.PostForm("http://localhost:8080/api/v1/socket", post)
	if err != nil {
		panic(err)
	}
}

func printHelp() {
	fmt.Println("Cli-2-Cloud")
	fmt.Println("A command line interface for piping output to web")
	exists, value := checkKey()
	if exists {
		fmt.Printf("Key: %s\n", value)
		fmt.Printf("Web-UI: https://cli2cloud.herokuapp.com/#/cli?id=%s\n", value)
	}
	fmt.Println("Version: 0.1.1")
	fmt.Println("Usage:")
	fmt.Println("\t$ command | c2c")
	fmt.Println("\t$ command | c2c q  quiet mode")
}

func main() {
	args := os.Args[1:]
	exists, value := checkKey()
	if exists {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			println(string(data))
			if len(args) == 0 {
				fmt.Printf("Web-UI: https://cli2cloud.herokuapp.com/#/cli?id=%s\n", value)
			}
			sendData(string(data), value)
		} else {
			printHelp()
			os.Exit(0)
		}
	} else {
		createKey()
	}
}

//usr/bin/go run $0$@ ; exit

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"github.com/jkomyno/nanoid"
)

func checkKey() (exists bool, value string) {
	id, err := ioutil.ReadFile("c2c.dat")
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
	id, err := nanoid.Nanoid(8)
	if err != nil {
		fmt.Println(err)
	}
	writeErr := ioutil.WriteFile("c2c.dat", []byte(id), 0777)
	if err != nil {
		fmt.Println(writeErr)
	}
	fmt.Printf("Created key: %s\n", id)
}

func sendData(data string, value string) {
	post := url.Values{
		"data": {data},
		"room": {value},
	}
	_, err := http.PostForm("https://cli2cloud.herokuapp.com/api/v1/socket", post)
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
	}
	fmt.Println("Version: 0.1.0")
	fmt.Println("Usage:")
	fmt.Println("\t$ command | c2c")
}

func main() {
	exists, value := checkKey()
	if exists {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
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

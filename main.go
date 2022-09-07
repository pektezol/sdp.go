package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bisaxa/demoparser/messages"
	"github.com/bisaxa/demoparser/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Specify file in command line arguments.")
	}
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil { // If it's not a directory
		file, err := os.Open(os.Args[1])
		utils.CheckError(err)
		messages.ParseHeader(file)
		for {
			code := messages.ParseMessage(file)
			if code == 7 {
				messages.ParseMessage(file)
				break
			}
		}
		defer file.Close()
	}
	for _, fileinfo := range files { // If it is a directory
		file, err := os.Open(os.Args[1] + fileinfo.Name())
		utils.CheckError(err)
		messages.ParseHeader(file)
		for {
			code := messages.ParseMessage(file)
			if code == 7 {
				messages.ParseMessage(file)
				break
			}
		}
		defer file.Close()
	}
	fmt.Scanln()
}

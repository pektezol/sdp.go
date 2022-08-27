package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"parser/messages"
	"parser/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Specify file in command line arguments.")
	}
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil { // If it's not a directory
		file, err := os.Open(os.Args[1])
		utils.CheckError(err)
		utils.HeaderOut(file)
		for {
			code := messages.MessageTypeCheck(file)
			if code == 7 {
				messages.MessageTypeCheck(file)
				break // TODO: Check last CustomData
			}
		}
		defer file.Close()
	}
	for _, fileinfo := range files { // If it is a directory
		file, err := os.Open(os.Args[1] + fileinfo.Name())
		utils.CheckError(err)
		utils.HeaderOut(file)
		for {
			code := messages.MessageTypeCheck(file)
			if code == 7 {
				messages.MessageTypeCheck(file)
				break // TODO: Check last CustomData
			}
		}
		defer file.Close()
	}

	fmt.Scanln()
}

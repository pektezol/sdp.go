package main

import (
	"log"
	"os"
	"parser/messages"
	"parser/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Specifiy file in command line arguments.")
	}
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

	//defer file.Close()
}

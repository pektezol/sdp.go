package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/packets"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Specify file in command line arguments.")
	}
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil { // If it's not a directory
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		reader := bitreader.Reader(file, true)
		packets.ParseHeader(reader)
		for {
			code := packets.ParsePacket(reader)
			if code == 7 {
				break
			}
		}
		defer file.Close()
	}
	for _, fileinfo := range files { // If it is a directory
		file, err := os.Open(os.Args[1] + fileinfo.Name())
		if err != nil {
			panic(err)
		}
		/*messages.ParseHeader(file)
		for {
			code := messages.ParseMessage(file)
			if code == 7 {
				messages.ParseMessage(file)
				break
			}
		}*/
		defer file.Close()
	}
	fmt.Scanln()
}

package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/packets"
)

func main() {
	if len(os.Args) != 2 {
		panic("specify file in command line arguments")
	}
	files, err := os.ReadDir(os.Args[1])
	if err != nil { // If it's not a directory
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		reader := bitreader.Reader(file, true)
		demoParserHandler(reader)
		defer file.Close()
	}
	for _, fileinfo := range files { // If it is a directory
		file, err := os.Open(os.Args[1] + fileinfo.Name())
		if err != nil {
			panic(err)
		}
		reader := bitreader.Reader(file, true)
		demoParserHandler(reader)
		defer file.Close()
	}
	// fmt.Scanln()
}

func demoParserHandler(reader *bitreader.ReaderType) {
	packets.ParseHeaders(reader)
	for {
		packet := packets.ParsePackets(reader)
		if packet.PacketType == 7 {
			break
		}
		// if packet.PacketType != 5 {
		// 	continue
		// }
		fmt.Printf("[%d] %s (%d):\n\t%+v\n", packet.TickNumber, reflect.ValueOf(packet.Data).Type(), packet.PacketType, packet.Data)
	}
}

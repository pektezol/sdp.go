package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/packets"
)

const littleEndian bool = true

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
		reader := bitreader.NewReader(file, littleEndian)
		demoParserHandler(reader)
		defer file.Close()
		return
	}
	for _, fileinfo := range files { // If it is a directory
		file, err := os.Open(os.Args[1] + fileinfo.Name())
		if err != nil {
			panic(err)
		}
		reader := bitreader.NewReader(file, littleEndian)
		demoParserHandler(reader)
		defer file.Close()
	}
}

func demoParserHandler(reader *bitreader.Reader) {
	packets.ParseHeaders(reader)
	for {
		packet := packets.ParsePackets(reader)
		fmt.Printf("[%d] %s (%d):\n\t%+v\n", packet.TickNumber, reflect.ValueOf(packet.Data).Type(), packet.PacketType, packet.Data)
		if packet.PacketType == 7 {
			break
		}
	}
}

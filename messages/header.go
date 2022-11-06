package messages

import (
	"fmt"
	"os"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/utils"
)

func ParseHeader(file *os.File) {
	var header Header
	reader := bitreader.Reader(file, true)
	header.DemoFileStamp = string(utils.ReadByteFromFile(file, 8))
	header.DemoProtocol = int32(reader.TryReadInt32())
	header.NetworkProtocol = int32(reader.TryReadInt32())
	header.ServerName = string(utils.ReadByteFromFile(file, 260))
	header.ClientName = string(utils.ReadByteFromFile(file, 260))
	header.MapName = string(utils.ReadByteFromFile(file, 260))
	header.GameDirectory = string(utils.ReadByteFromFile(file, 260))
	header.PlaybackTime = float32(reader.TryReadFloat32())
	header.PlaybackTicks = int32(reader.TryReadInt32())
	header.PlaybackFrames = int32(reader.TryReadInt32())
	header.SignOnLength = int32(reader.TryReadInt32())
	if header.DemoProtocol != 4 { // Old Engine == 2, 3 / New Engine == 4
		panic("Only New Engine is supported.")
	}
	fmt.Printf("%+v", header)
}

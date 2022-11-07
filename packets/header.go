package packets

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

func ParseHeader(reader *bitreader.ReaderType) {
	header := Header{
		DemoFileStamp:   reader.TryReadStringLen(8),
		DemoProtocol:    uint32(reader.TryReadInt32()),
		NetworkProtocol: uint32(reader.TryReadInt32()),
		ServerName:      reader.TryReadStringLen(260),
		ClientName:      reader.TryReadStringLen(260),
		MapName:         reader.TryReadStringLen(260),
		GameDirectory:   reader.TryReadStringLen(260),
		PlaybackTime:    reader.TryReadFloat32(),
		PlaybackTicks:   int32(reader.TryReadInt32()),
		PlaybackFrames:  int32(reader.TryReadInt32()),
		SignOnLength:    uint32(reader.TryReadInt32()),
	}
	if header.DemoFileStamp != "HL2DEMO" {
		panic("Invalid demo file stamp. Make sure a valid demo file is given.")
	}
	if header.DemoProtocol != 4 {
		panic("Given demo is from old engine. This parser is only supported for new engine.")
	}
	if header.NetworkProtocol != 2001 {
		panic("Given demo is not from Portal2. This parser currently only supports Portal 2.")
	}
	fmt.Println(header)
}

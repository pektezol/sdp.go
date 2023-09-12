package packets

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type Headers struct {
	DemoFileStamp   string
	DemoProtocol    int32
	NetworkProtocol int32
	ServerName      string
	ClientName      string
	MapName         string
	GameDirectory   string
	PlaybackTime    float32
	PlaybackTicks   int32
	PlaybackFrames  int32
	SignOnLength    int32
}

func ParseHeaders(reader *bitreader.ReaderType) Headers {
	headers := Headers{
		DemoFileStamp:   reader.TryReadString(),
		DemoProtocol:    int32(reader.TryReadInt32()),
		NetworkProtocol: int32(reader.TryReadInt32()),
		ServerName:      reader.TryReadStringLen(260),
		ClientName:      reader.TryReadStringLen(260),
		MapName:         reader.TryReadStringLen(260),
		GameDirectory:   reader.TryReadStringLen(260),
		PlaybackTime:    reader.TryReadFloat32(),
		PlaybackTicks:   int32(reader.TryReadInt32()),
		PlaybackFrames:  int32(reader.TryReadInt32()),
		SignOnLength:    int32(reader.TryReadInt32()),
	}
	if headers.DemoFileStamp != "HL2DEMO" {
		panic("invalid demo file stamp")
	}
	if headers.DemoProtocol != 4 {
		panic("this parser only supports demos from new engine")
	}
	if headers.NetworkProtocol != 2001 {
		panic("this parser only supports demos from portal 2")
	}
	fmt.Printf("Headers: %+v\n", headers)
	return headers
}

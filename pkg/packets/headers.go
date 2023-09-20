package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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

func ParseHeaders(reader *bitreader.Reader) Headers {
	headers := Headers{
		DemoFileStamp:   reader.TryReadString(),
		DemoProtocol:    int32(reader.TryReadSInt32()),
		NetworkProtocol: int32(reader.TryReadSInt32()),
		ServerName:      reader.TryReadStringLength(260),
		ClientName:      reader.TryReadStringLength(260),
		MapName:         reader.TryReadStringLength(260),
		GameDirectory:   reader.TryReadStringLength(260),
		PlaybackTime:    reader.TryReadFloat32(),
		PlaybackTicks:   int32(reader.TryReadSInt32()),
		PlaybackFrames:  int32(reader.TryReadSInt32()),
		SignOnLength:    int32(reader.TryReadSInt32()),
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
	writer.AppendLine("Headers: %+v", headers)
	return headers
}

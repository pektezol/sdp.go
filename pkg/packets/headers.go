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
	writer.AppendLine("\nDemo File Stamp: %s", headers.DemoFileStamp)
	writer.AppendLine("Demo Protocol: %d", headers.DemoProtocol)
	writer.AppendLine("Network Protocol: %d", headers.NetworkProtocol)
	writer.AppendLine("Server Name: %s", headers.ServerName)
	writer.AppendLine("Client Name: %s", headers.ClientName)
	writer.AppendLine("Map Name: %s", headers.MapName)
	writer.AppendLine("Game Directory: %s", headers.GameDirectory)
	writer.AppendLine("Playback Time: %f", headers.PlaybackTime)
	writer.AppendLine("Playback Ticks: %d", headers.PlaybackTicks)
	writer.AppendLine("Playback Frames: %d", headers.PlaybackFrames)
	writer.AppendLine("Sign On Length: %d\n", headers.SignOnLength)
	return headers
}

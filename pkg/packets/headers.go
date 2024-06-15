package packets

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

func ParseHeaders(reader *bitreader.Reader, demo *types.Demo) types.Headers {
	headers := types.Headers{
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
	demo.Writer.AppendLine("\nDemo File Stamp: %s", headers.DemoFileStamp)
	demo.Writer.AppendLine("Demo Protocol: %d", headers.DemoProtocol)
	demo.Writer.AppendLine("Network Protocol: %d", headers.NetworkProtocol)
	demo.Writer.AppendLine("Server Name: %s", headers.ServerName)
	demo.Writer.AppendLine("Client Name: %s", headers.ClientName)
	demo.Writer.AppendLine("Map Name: %s", headers.MapName)
	demo.Writer.AppendLine("Game Directory: %s", headers.GameDirectory)
	demo.Writer.AppendLine("Playback Time: %f", headers.PlaybackTime)
	demo.Writer.AppendLine("Playback Ticks: %d", headers.PlaybackTicks)
	demo.Writer.AppendLine("Playback Frames: %d", headers.PlaybackFrames)
	demo.Writer.AppendLine("Sign On Length: %d\n", headers.SignOnLength)
	return headers
}

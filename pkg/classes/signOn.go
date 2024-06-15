package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/messages"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SignOn struct {
	PacketInfo  []CmdInfo `json:"packet_info"`
	InSequence  uint32    `json:"in_sequence"`
	OutSequence uint32    `json:"out_sequence"`
	Size        uint32    `json:"size"`
	Data        []any     `json:"data"`
}

func (signOn *SignOn) ParseSignOn(reader *bitreader.Reader, demo *types.Demo) {
	for count := 0; count < MSSC; count++ {
		signOn.ParseCmdInfo(reader, demo)
	}
	signOn.InSequence = reader.TryReadUInt32()
	signOn.OutSequence = reader.TryReadUInt32()
	signOn.Size = reader.TryReadUInt32()
	packetReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(signOn.Size)), true)
	for {
		messageType, err := packetReader.ReadBits(6)
		if err != nil {
			break
		}
		signOn.Data = append(signOn.Data, messages.ParseMessages(messageType, packetReader, demo))
	}
}

func (signOn *SignOn) ParseCmdInfo(reader *bitreader.Reader, demo *types.Demo) {
	signOn.PacketInfo = append(signOn.PacketInfo, CmdInfo{
		Flags:            reader.TryReadUInt32(),
		ViewOrigin:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewOrigin2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles2: []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	})
	demo.Writer.AppendLine("\tFlags: %s", CmdInfoFlags(signOn.PacketInfo[len(signOn.PacketInfo)-1].Flags).String())
	demo.Writer.AppendLine("\tView Origin: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewOrigin)
	demo.Writer.AppendLine("\tView Angles: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewAngles)
	demo.Writer.AppendLine("\tLocal View Angles: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].LocalViewAngles)
	demo.Writer.AppendLine("\tView Origin 2: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewOrigin2)
	demo.Writer.AppendLine("\tView Angles 2: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewAngles2)
	demo.Writer.AppendLine("\tLocal View Angles 2: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].LocalViewAngles2)
	demo.Writer.AppendLine("")
}

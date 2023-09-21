package classes

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/messages"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SignOn struct {
	PacketInfo  []CmdInfo
	InSequence  uint32
	OutSequence uint32
	Size        uint32
	Data        []any
}

func (signOn *SignOn) ParseSignOn(reader *bitreader.Reader) {
	for count := 0; count < MSSC; count++ {
		signOn.ParseCmdInfo(reader)
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
		signOn.Data = append(signOn.Data, messages.ParseMessages(messageType, packetReader))
	}
}

func (signOn *SignOn) ParseCmdInfo(reader *bitreader.Reader) {
	signOn.PacketInfo = append(signOn.PacketInfo, CmdInfo{
		Flags:            reader.TryReadUInt32(),
		ViewOrigin:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewOrigin2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles2: []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	})
	writer.AppendLine("\tFlags: %s", CmdInfoFlags(signOn.PacketInfo[len(signOn.PacketInfo)-1].Flags).String())
	writer.AppendLine("\tView Origin: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewOrigin)
	writer.AppendLine("\tView Angles: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewAngles)
	writer.AppendLine("\tLocal View Angles: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].LocalViewAngles)
	writer.AppendLine("\tView Origin 2: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewOrigin2)
	writer.AppendLine("\tView Angles 2: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].ViewAngles2)
	writer.AppendLine("\tLocal View Angles 2: %v", signOn.PacketInfo[len(signOn.PacketInfo)-1].LocalViewAngles2)
	writer.AppendLine("")
}

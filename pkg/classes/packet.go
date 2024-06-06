package classes

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/messages"
	"github.com/pektezol/sdp.go/pkg/writer"
)

const MSSC int = 2

type Packet struct {
	PacketInfo  []CmdInfo
	InSequence  uint32
	OutSequence uint32
	Size        uint32
	Data        []any
}

type CmdInfo struct {
	Flags            uint32
	ViewOrigin       []float32
	ViewAngles       []float32
	LocalViewAngles  []float32
	ViewOrigin2      []float32
	ViewAngles2      []float32
	LocalViewAngles2 []float32
}

func (packet *Packet) ParsePacket(reader *bitreader.Reader) {
	for count := 0; count < MSSC; count++ {
		packet.ParseCmdInfo(reader)
	}
	packet.InSequence = reader.TryReadUInt32()
	packet.OutSequence = reader.TryReadUInt32()
	packet.Size = reader.TryReadUInt32()
	packetReader := bitreader.NewReaderFromBytes(reader.TryReadBytesToSlice(uint64(packet.Size)), true)
	for {
		messageType, err := packetReader.ReadBits(6)
		if err != nil {
			break
		}
		packet.Data = append(packet.Data, messages.ParseMessages(messageType, packetReader))
	}
}

func (packet *Packet) ParseCmdInfo(reader *bitreader.Reader) {
	packet.PacketInfo = append(packet.PacketInfo, CmdInfo{
		Flags:            reader.TryReadUInt32(),
		ViewOrigin:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewOrigin2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles2: []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	})
	writer.AppendLine("\tFlags: %s", CmdInfoFlags(packet.PacketInfo[len(packet.PacketInfo)-1].Flags).String())
	writer.AppendLine("\tView Origin: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewOrigin)
	writer.AppendLine("\tView Angles: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewAngles)
	writer.AppendLine("\tLocal View Angles: %v", packet.PacketInfo[len(packet.PacketInfo)-1].LocalViewAngles)
	writer.AppendLine("\tView Origin 2: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewOrigin2)
	writer.AppendLine("\tView Angles 2: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewAngles2)
	writer.AppendLine("\tLocal View Angles 2: %v", packet.PacketInfo[len(packet.PacketInfo)-1].LocalViewAngles2)
	writer.AppendLine("")
}

type CmdInfoFlags int

const (
	ECmdInfoFlagsNone        = 0
	ECmdInfoFlagsUseOrigin2  = 1
	ECmdInfoFlagsUserAngles2 = 1 << 1
	ECmdInfoFlagsNoInterp    = 1 << 2
)

func (cmdInfoFlags CmdInfoFlags) String() string {
	switch cmdInfoFlags {
	case ECmdInfoFlagsNone:
		return "None"
	case ECmdInfoFlagsUseOrigin2:
		return "UseOrigin2"
	case ECmdInfoFlagsUserAngles2:
		return "UserAngles2"
	case ECmdInfoFlagsNoInterp:
		return "NoInterp"
	default:
		return fmt.Sprintf("%d", int(cmdInfoFlags))
	}
}

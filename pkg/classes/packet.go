package classes

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/messages"
	"github.com/pektezol/sdp.go/pkg/types"
)

const MSSC int = 2

type Packet struct {
	PacketInfo  []CmdInfo `json:"packet_info"`
	InSequence  uint32    `json:"in_sequence"`
	OutSequence uint32    `json:"out_sequence"`
	Size        uint32    `json:"size"`
	Data        []any     `json:"data"`
}

type CmdInfo struct {
	Flags            uint32    `json:"flags"`
	ViewOrigin       []float32 `json:"view_origin"`
	ViewAngles       []float32 `json:"view_angles"`
	LocalViewAngles  []float32 `json:"local_view_angles"`
	ViewOrigin2      []float32 `json:"view_origin_2"`
	ViewAngles2      []float32 `json:"view_angles_2"`
	LocalViewAngles2 []float32 `json:"local_view_angles_2"`
}

func (packet *Packet) ParsePacket(reader *bitreader.Reader, demo *types.Demo) {
	for count := 0; count < MSSC; count++ {
		packet.ParseCmdInfo(reader, demo)
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
		packet.Data = append(packet.Data, messages.ParseMessages(messageType, packetReader, demo))
	}
}

func (packet *Packet) ParseCmdInfo(reader *bitreader.Reader, demo *types.Demo) {
	packet.PacketInfo = append(packet.PacketInfo, CmdInfo{
		Flags:            reader.TryReadUInt32(),
		ViewOrigin:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles:       []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles:  []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewOrigin2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		ViewAngles2:      []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
		LocalViewAngles2: []float32{reader.TryReadFloat32(), reader.TryReadFloat32(), reader.TryReadFloat32()},
	})
	demo.Writer.AppendLine("\tFlags: %s", CmdInfoFlags(packet.PacketInfo[len(packet.PacketInfo)-1].Flags).String())
	demo.Writer.AppendLine("\tView Origin: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewOrigin)
	demo.Writer.AppendLine("\tView Angles: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewAngles)
	demo.Writer.AppendLine("\tLocal View Angles: %v", packet.PacketInfo[len(packet.PacketInfo)-1].LocalViewAngles)
	demo.Writer.AppendLine("\tView Origin 2: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewOrigin2)
	demo.Writer.AppendLine("\tView Angles 2: %v", packet.PacketInfo[len(packet.PacketInfo)-1].ViewAngles2)
	demo.Writer.AppendLine("\tLocal View Angles 2: %v", packet.PacketInfo[len(packet.PacketInfo)-1].LocalViewAngles2)
	demo.Writer.AppendLine("")
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

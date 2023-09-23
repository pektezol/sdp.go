package classes

import (
	"github.com/pektezol/bitreader"
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
		reader.SkipBytes(76)
	}
	reader.SkipBytes(8)
	reader.TryReadBytesToSlice(uint64(reader.TryReadUInt32()))
}

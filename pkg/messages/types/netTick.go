package messages

import (
	"github.com/pektezol/bitreader"
)

type NetTick struct {
	Tick                      uint32
	HostFrameTime             float32
	HostFrameTimeStdDeviation float32
}

func ParseNetTick(reader *bitreader.Reader) NetTick {
	netTick := NetTick{
		Tick:                      reader.TryReadUInt32(),
		HostFrameTime:             float32(reader.TryReadUInt16()) / 1e5,
		HostFrameTimeStdDeviation: float32(reader.TryReadUInt16()) / 1e5,
	}

	return netTick
}

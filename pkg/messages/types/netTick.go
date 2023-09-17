package messages

import "github.com/pektezol/bitreader"

type NetTick struct {
	Tick                      int32
	HostFrameTime             float32
	HostFrameTimeStdDeviation float32
}

func ParseNetTick(reader *bitreader.Reader) NetTick {
	return NetTick{
		Tick:                      int32(reader.TryReadBits(32)),
		HostFrameTime:             float32(reader.TryReadBits(16)) / 1e5,
		HostFrameTimeStdDeviation: float32(reader.TryReadBits(16)) / 1e5,
	}
}

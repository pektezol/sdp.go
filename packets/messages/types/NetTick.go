package types

import "github.com/pektezol/bitreader"

type NetTick struct {
	Tick                      uint32
	HostFrameTime             float32
	HostFrameTimeStdDeviation float32
}

func ParseNetTick(reader *bitreader.ReaderType) NetTick {
	return NetTick{
		Tick:                      reader.TryReadInt32(),
		HostFrameTime:             float32(reader.TryReadInt16()) / 1e5,
		HostFrameTimeStdDeviation: float32(reader.TryReadInt16()) / 1e5,
	}
}

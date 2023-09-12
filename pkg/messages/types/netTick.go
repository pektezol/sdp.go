package messages

import "github.com/pektezol/bitreader"

type NetTick struct {
	Tick                      int32
	HostFrameTime             int16
	HostFrameTimeStdDeviation int16
}

func ParseNetTick(reader *bitreader.ReaderType) NetTick {
	return NetTick{
		Tick:                      int32(reader.TryReadBits(32)),
		HostFrameTime:             int16(reader.TryReadBits(16) / 10e5),
		HostFrameTimeStdDeviation: int16(reader.TryReadBits(16) / 10e5),
	}
}

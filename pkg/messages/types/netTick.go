package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	writer.TempAppendLine("\t\tTick: %d", netTick.Tick)
	writer.TempAppendLine("\t\tHost Frame Time: %f", netTick.HostFrameTime)
	writer.TempAppendLine("\t\tHost Frame Time Std Deviation: %f", netTick.HostFrameTimeStdDeviation)
	return netTick
}

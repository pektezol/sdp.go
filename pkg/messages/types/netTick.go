package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetTick struct {
	Tick                      uint32  `json:"tick"`
	HostFrameTime             float32 `json:"host_frame_time"`
	HostFrameTimeStdDeviation float32 `json:"host_frame_time_std_deviation"`
}

func ParseNetTick(reader *bitreader.Reader, demo *types.Demo) NetTick {
	netTick := NetTick{
		Tick:                      reader.TryReadUInt32(),
		HostFrameTime:             float32(reader.TryReadUInt16()) / 1e5,
		HostFrameTimeStdDeviation: float32(reader.TryReadUInt16()) / 1e5,
	}
	demo.Writer.TempAppendLine("\t\tTick: %d", netTick.Tick)
	demo.Writer.TempAppendLine("\t\tHost Frame Time: %f", netTick.HostFrameTime)
	demo.Writer.TempAppendLine("\t\tHost Frame Time Std Deviation: %f", netTick.HostFrameTimeStdDeviation)
	return netTick
}

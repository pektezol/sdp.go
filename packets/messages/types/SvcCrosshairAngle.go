package types

import "github.com/pektezol/bitreader"

type SvcCrosshairAngle struct {
	Angle []int16
}

func ParseSvcCrosshairAngle(reader *bitreader.ReaderType) SvcCrosshairAngle {
	return SvcCrosshairAngle{
		Angle: []int16{
			int16(reader.TryReadInt16()),
			int16(reader.TryReadInt16()),
			int16(reader.TryReadInt16()),
		},
	}
}

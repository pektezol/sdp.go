package messages

import "github.com/pektezol/bitreader"

type SvcCrosshairAngle struct {
	Angle []int16
}

func ParseSvcCrosshairAngle(reader *bitreader.Reader) SvcCrosshairAngle {
	return SvcCrosshairAngle{
		Angle: []int16{int16(reader.TryReadBits(16)), int16(reader.TryReadBits(16)), int16(reader.TryReadBits(16))},
	}
}

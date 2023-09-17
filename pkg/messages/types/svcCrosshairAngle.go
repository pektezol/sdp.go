package messages

import "github.com/pektezol/bitreader"

type SvcCrosshairAngle struct {
	Angle crosshairAngles
}

type crosshairAngles struct {
	X float32
	Y float32
	Z float32
}

func ParseSvcCrosshairAngle(reader *bitreader.Reader) SvcCrosshairAngle {
	return SvcCrosshairAngle{
		Angle: crosshairAngles{
			X: float32(reader.TryReadBits(16)),
			Y: float32(reader.TryReadBits(16)),
			Z: float32(reader.TryReadBits(16)),
		},
	}
}

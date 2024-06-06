package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type SvcCrosshairAngle struct {
	Angle crosshairAngles
}

type crosshairAngles struct {
	X float32
	Y float32
	Z float32
}

func ParseSvcCrosshairAngle(reader *bitreader.Reader) SvcCrosshairAngle {
	svcCrosshairAngle := SvcCrosshairAngle{
		Angle: crosshairAngles{
			X: float32(reader.TryReadBits(16)),
			Y: float32(reader.TryReadBits(16)),
			Z: float32(reader.TryReadBits(16)),
		},
	}
	writer.TempAppendLine("\t\tX: %f", svcCrosshairAngle.Angle.X)
	writer.TempAppendLine("\t\tY: %f", svcCrosshairAngle.Angle.Y)
	writer.TempAppendLine("\t\tZ: %f", svcCrosshairAngle.Angle.Z)
	return svcCrosshairAngle
}

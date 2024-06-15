package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcCrosshairAngle struct {
	Angle crosshairAngles `json:"angle"`
}

type crosshairAngles struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func ParseSvcCrosshairAngle(reader *bitreader.Reader, demo *types.Demo) SvcCrosshairAngle {
	svcCrosshairAngle := SvcCrosshairAngle{
		Angle: crosshairAngles{
			X: float32(reader.TryReadBits(16)),
			Y: float32(reader.TryReadBits(16)),
			Z: float32(reader.TryReadBits(16)),
		},
	}
	demo.Writer.TempAppendLine("\t\tX: %f", svcCrosshairAngle.Angle.X)
	demo.Writer.TempAppendLine("\t\tY: %f", svcCrosshairAngle.Angle.Y)
	demo.Writer.TempAppendLine("\t\tZ: %f", svcCrosshairAngle.Angle.Z)
	return svcCrosshairAngle
}

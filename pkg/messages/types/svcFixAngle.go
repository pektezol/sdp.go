package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcFixAngle struct {
	Relative bool      `json:"relative"`
	Angle    fixAngles `json:"angle"`
}

type fixAngles struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func ParseSvcFixAngle(reader *bitreader.Reader, demo *types.Demo) SvcFixAngle {
	svcFixAngle := SvcFixAngle{
		Relative: reader.TryReadBool(),
		Angle: fixAngles{
			X: float32(reader.TryReadBits(16)),
			Y: float32(reader.TryReadBits(16)),
			Z: float32(reader.TryReadBits(16)),
		},
	}
	demo.Writer.TempAppendLine("\t\tRelative: %t", svcFixAngle.Relative)
	demo.Writer.TempAppendLine("\t\tX: %f", svcFixAngle.Angle.X)
	demo.Writer.TempAppendLine("\t\tY: %f", svcFixAngle.Angle.Y)
	demo.Writer.TempAppendLine("\t\tZ: %f", svcFixAngle.Angle.Z)
	return svcFixAngle
}

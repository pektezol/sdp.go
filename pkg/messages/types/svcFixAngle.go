package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type SvcFixAngle struct {
	Relative bool
	Angle    fixAngles
}

type fixAngles struct {
	X float32
	Y float32
	Z float32
}

func ParseSvcFixAngle(reader *bitreader.Reader) SvcFixAngle {
	svcFixAngle := SvcFixAngle{
		Relative: reader.TryReadBool(),
		Angle: fixAngles{
			X: float32(reader.TryReadBits(16)),
			Y: float32(reader.TryReadBits(16)),
			Z: float32(reader.TryReadBits(16)),
		},
	}
	writer.TempAppendLine("\t\tRelative: %t", svcFixAngle.Relative)
	writer.TempAppendLine("\t\tX: %f", svcFixAngle.Angle.X)
	writer.TempAppendLine("\t\tY: %f", svcFixAngle.Angle.Y)
	writer.TempAppendLine("\t\tZ: %f", svcFixAngle.Angle.Z)
	return svcFixAngle
}

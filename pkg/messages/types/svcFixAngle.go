package messages

import "github.com/pektezol/bitreader"

type SvcFixAngle struct {
	Relative bool
	Angle    []int16
}

func ParseSvcFixAngle(reader *bitreader.Reader) SvcFixAngle {
	return SvcFixAngle{
		Relative: reader.TryReadBool(),
		Angle:    []int16{int16(reader.TryReadBits(16)), int16(reader.TryReadBits(16)), int16(reader.TryReadBits(16))},
	}
}

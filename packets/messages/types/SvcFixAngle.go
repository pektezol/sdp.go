package types

import "github.com/pektezol/bitreader"

type SvcFixAngle struct {
	Relative bool
	Angle    []int16
}

func ParseSvcFixAngle(reader *bitreader.ReaderType) SvcFixAngle {
	relative := reader.TryReadBool()
	angles := []int16{
		int16(reader.TryReadInt16()),
		int16(reader.TryReadInt16()),
		int16(reader.TryReadInt16()),
	}
	return SvcFixAngle{
		Relative: relative,
		Angle:    angles,
	}
}

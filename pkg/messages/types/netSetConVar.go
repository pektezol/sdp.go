package messages

import (
	"github.com/pektezol/bitreader"
)

type NetSetConVar struct {
	Length  uint8
	ConVars []conVar
}

type conVar struct {
	Name  string
	Value string
}

func ParseNetSetConVar(reader *bitreader.Reader) NetSetConVar {
	length := reader.TryReadUInt8()
	convars := []conVar{}

	for count := 0; count < int(length); count++ {
		convar := conVar{
			Name:  reader.TryReadString(),
			Value: reader.TryReadString(),
		}

		convars = append(convars, convar)
	}
	return NetSetConVar{
		Length:  length,
		ConVars: convars,
	}
}

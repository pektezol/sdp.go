package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetSetConVar struct {
	Length  uint8    `json:"length"`
	ConVars []conVar `json:"con_vars"`
}

type conVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func ParseNetSetConVar(reader *bitreader.Reader, demo *types.Demo) NetSetConVar {
	length := reader.TryReadUInt8()
	convars := []conVar{}
	demo.Writer.TempAppendLine("\t\tLength: %d", length)
	for count := 0; count < int(length); count++ {
		convar := conVar{
			Name:  reader.TryReadString(),
			Value: reader.TryReadString(),
		}
		demo.Writer.TempAppendLine("\t\t[%d] %s: %s", count, convar.Name, convar.Value)
		convars = append(convars, convar)
	}
	return NetSetConVar{
		Length:  length,
		ConVars: convars,
	}
}

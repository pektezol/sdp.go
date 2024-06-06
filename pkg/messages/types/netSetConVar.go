package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
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
	writer.TempAppendLine("\t\tLength: %d", length)
	for count := 0; count < int(length); count++ {
		convar := conVar{
			Name:  reader.TryReadString(),
			Value: reader.TryReadString(),
		}
		writer.TempAppendLine("\t\t[%d] %s: %s", count, convar.Name, convar.Value)
		convars = append(convars, convar)
	}
	return NetSetConVar{
		Length:  length,
		ConVars: convars,
	}
}

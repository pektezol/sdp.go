package types

import "github.com/pektezol/bitreader"

type NetSetConVar struct {
	ConVars []ConVar
}

type ConVar struct {
	Name  string
	Value string
}

func ParseNetSetConVar(reader *bitreader.ReaderType) NetSetConVar {
	length := reader.TryReadInt8()
	convars := make([]ConVar, length)
	for i := 0; i < int(length); i++ {
		convars[i] = ConVar{
			Name:  reader.TryReadString(),
			Value: reader.TryReadString(),
		}
	}
	return NetSetConVar{
		ConVars: convars,
	}
}

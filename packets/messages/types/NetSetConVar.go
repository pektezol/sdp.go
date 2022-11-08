package types

import "github.com/pektezol/bitreader"

type NetSetConVar struct {
	Length  uint8
	ConVars []ConVar
}

type ConVar struct {
	Name  string
	Value string
}

func ParseNetSetConVar(reader *bitreader.ReaderType) NetSetConVar {
	var convars []ConVar
	netsetconvar := NetSetConVar{
		Length: reader.TryReadInt8(),
	}
	for i := 0; i < int(netsetconvar.Length); i++ {
		convars = append(convars, ConVar{
			Name:  reader.TryReadString(),
			Value: reader.TryReadString(),
		})
	}
	netsetconvar.ConVars = convars
	return netsetconvar
}

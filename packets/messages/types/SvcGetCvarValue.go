package types

import "github.com/pektezol/bitreader"

type SvcGetCvarValue struct {
	Cookie   string
	CvarName string
}

func ParseSvcGetCvarValue(reader *bitreader.ReaderType) SvcGetCvarValue {
	return SvcGetCvarValue{
		Cookie:   reader.TryReadStringLen(32),
		CvarName: reader.TryReadString(),
	}
}

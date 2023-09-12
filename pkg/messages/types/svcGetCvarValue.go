package messages

import "github.com/pektezol/bitreader"

type SvcGetCvarValue struct {
	Cookie   string
	CvarName string
}

func ParseSvcGetCvarValue(reader *bitreader.ReaderType) SvcGetCvarValue {
	svcGetCvarValue := SvcGetCvarValue{
		Cookie:   reader.TryReadStringLen(4),
		CvarName: reader.TryReadString(),
	}
	return svcGetCvarValue
}

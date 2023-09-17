package messages

import "github.com/pektezol/bitreader"

type SvcGetCvarValue struct {
	Cookie   int32
	CvarName string
}

func ParseSvcGetCvarValue(reader *bitreader.Reader) SvcGetCvarValue {
	svcGetCvarValue := SvcGetCvarValue{
		Cookie:   reader.TryReadSInt32(),
		CvarName: reader.TryReadString(),
	}
	return svcGetCvarValue
}

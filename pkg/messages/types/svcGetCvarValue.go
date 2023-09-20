package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcGetCvarValue struct {
	Cookie   int32
	CvarName string
}

func ParseSvcGetCvarValue(reader *bitreader.Reader) SvcGetCvarValue {
	svcGetCvarValue := SvcGetCvarValue{
		Cookie:   reader.TryReadSInt32(),
		CvarName: reader.TryReadString(),
	}
	writer.TempAppendLine("\t\tCookie: %d", svcGetCvarValue.Cookie)
	writer.TempAppendLine("\t\tCvar: \"%s\"", svcGetCvarValue.CvarName)
	return svcGetCvarValue
}

package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcGetCvarValue struct {
	Cookie   int32  `json:"cookie"`
	CvarName string `json:"cvar_name"`
}

func ParseSvcGetCvarValue(reader *bitreader.Reader, demo *types.Demo) SvcGetCvarValue {
	svcGetCvarValue := SvcGetCvarValue{
		Cookie:   reader.TryReadSInt32(),
		CvarName: reader.TryReadString(),
	}
	demo.Writer.TempAppendLine("\t\tCookie: %d", svcGetCvarValue.Cookie)
	demo.Writer.TempAppendLine("\t\tCvar: \"%s\"", svcGetCvarValue.CvarName)
	return svcGetCvarValue
}

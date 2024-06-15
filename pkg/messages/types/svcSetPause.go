package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcSetPause struct {
	Paused bool `json:"paused"`
}

func ParseSvcSetPause(reader *bitreader.Reader, demo *types.Demo) SvcSetPause {
	svcSetPause := SvcSetPause{
		Paused: reader.TryReadBool(),
	}
	demo.Writer.TempAppendLine("\t\tPaused: %t", svcSetPause.Paused)
	return svcSetPause
}

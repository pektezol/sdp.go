package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcSetPause struct {
	Paused bool
}

func ParseSvcSetPause(reader *bitreader.Reader) SvcSetPause {
	svcSetPause := SvcSetPause{
		Paused: reader.TryReadBool(),
	}
	writer.TempAppendLine("\t\tPaused: %t", svcSetPause.Paused)
	return svcSetPause
}

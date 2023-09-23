package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcSetPause struct {
	Paused bool
}

func ParseSvcSetPause(reader *bitreader.Reader) SvcSetPause {
	svcSetPause := SvcSetPause{
		Paused: reader.TryReadBool(),
	}

	return svcSetPause
}

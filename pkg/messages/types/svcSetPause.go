package messages

import "github.com/pektezol/bitreader"

type SvcSetPause struct {
	Paused bool
}

func ParseSvcSetPause(reader *bitreader.Reader) SvcSetPause {
	return SvcSetPause{
		Paused: reader.TryReadBool(),
	}
}

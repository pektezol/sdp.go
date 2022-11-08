package types

import "github.com/pektezol/bitreader"

type SvcSetPause struct {
	Paused bool
}

func ParseSvcSetPause(reader *bitreader.ReaderType) SvcSetPause {
	return SvcSetPause{Paused: reader.TryReadBool()}
}

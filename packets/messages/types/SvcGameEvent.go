package types

import (
	"github.com/pektezol/bitreader"
)

type SvcGameEvent struct {
	Data []byte
}

func ParseSvcGameEvent(reader *bitreader.ReaderType) SvcGameEvent {
	length := reader.TryReadBits(11)
	reader.SkipBits(int(length))
	return SvcGameEvent{} // TODO: Parse SvcGameEvent
}

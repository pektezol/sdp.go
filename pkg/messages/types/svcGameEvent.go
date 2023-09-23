package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcGameEvent struct {
	Length uint16
	Data   []byte // TODO: GameEvent[]
}

func ParseSvcGameEvent(reader *bitreader.Reader) SvcGameEvent {
	svcGameEvent := SvcGameEvent{
		Length: uint16(reader.TryReadBits(11)),
	}
	svcGameEvent.Data = reader.TryReadBitsToSlice(uint64(svcGameEvent.Length))

	return svcGameEvent
}

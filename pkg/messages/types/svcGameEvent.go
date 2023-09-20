package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	writer.TempAppendLine("\t\tData: %v", svcGameEvent.Data)
	return svcGameEvent
}

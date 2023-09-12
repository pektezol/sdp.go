package messages

import "github.com/pektezol/bitreader"

type SvcGameEvent struct {
	Length int16
	Data   []byte // TODO: GameEvent[]
}

func ParseSvcGameEvent(reader *bitreader.ReaderType) SvcGameEvent {
	svcGameEvent := SvcGameEvent{
		Length: int16(reader.TryReadBits(11)),
	}
	svcGameEvent.Data = reader.TryReadBitsToSlice(int(svcGameEvent.Length))
	return svcGameEvent
}

package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcGameEventList struct {
	Events              int16
	Length              int32
	GameEventDescriptor []gameEventDescriptor
}

type gameEventDescriptor struct {
}

func ParseSvcGameEventList(reader *bitreader.Reader) SvcGameEventList {
	svcGameEventList := SvcGameEventList{
		Events: int16(reader.TryReadBits(9)),
		Length: int32(reader.TryReadBits(20)),
	}
	reader.TryReadBitsToSlice(uint64(svcGameEventList.Length))
	writer.TempAppendLine("\t\tEvents: %d", svcGameEventList.Events)
	return svcGameEventList
}

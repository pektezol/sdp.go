package messages

import "github.com/pektezol/bitreader"

type SvcEntityMessage struct {
	EntityIndex int16
	ClassId     int16
	Length      int16
	Data        []byte
}

func ParseSvcEntityMessage(reader *bitreader.ReaderType) SvcEntityMessage {
	svcEntityMessage := SvcEntityMessage{
		EntityIndex: int16(reader.TryReadBits(11)),
		ClassId:     int16(reader.TryReadBits(9)),
		Length:      int16(reader.TryReadBits(11)),
	}
	svcEntityMessage.Data = reader.TryReadBitsToSlice(int(svcEntityMessage.Length))
	return svcEntityMessage
}

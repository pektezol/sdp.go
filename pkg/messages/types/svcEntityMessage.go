package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type SvcEntityMessage struct {
	EntityIndex uint16
	ClassId     uint16
	Length      uint16
	Data        []byte
}

func ParseSvcEntityMessage(reader *bitreader.Reader) SvcEntityMessage {
	svcEntityMessage := SvcEntityMessage{
		EntityIndex: uint16(reader.TryReadBits(11)),
		ClassId:     uint16(reader.TryReadBits(9)),
		Length:      uint16(reader.TryReadBits(11)),
	}
	svcEntityMessage.Data = reader.TryReadBitsToSlice(uint64(svcEntityMessage.Length))
	writer.TempAppendLine("\t\tEntity Index: %d", svcEntityMessage.EntityIndex)
	writer.TempAppendLine("\t\tClass ID: %d", svcEntityMessage.ClassId)
	writer.TempAppendLine("\t\tData: %v", svcEntityMessage.Data)
	return svcEntityMessage
}

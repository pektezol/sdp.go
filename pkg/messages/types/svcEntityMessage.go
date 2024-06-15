package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcEntityMessage struct {
	EntityIndex uint16 `json:"entity_index"`
	ClassId     uint16 `json:"class_id"`
	Length      uint16 `json:"length"`
	Data        []byte `json:"data"`
}

func ParseSvcEntityMessage(reader *bitreader.Reader, demo *types.Demo) SvcEntityMessage {
	svcEntityMessage := SvcEntityMessage{
		EntityIndex: uint16(reader.TryReadBits(11)),
		ClassId:     uint16(reader.TryReadBits(9)),
		Length:      uint16(reader.TryReadBits(11)),
	}
	svcEntityMessage.Data = reader.TryReadBitsToSlice(uint64(svcEntityMessage.Length))
	demo.Writer.TempAppendLine("\t\tEntity Index: %d", svcEntityMessage.EntityIndex)
	demo.Writer.TempAppendLine("\t\tClass ID: %d", svcEntityMessage.ClassId)
	demo.Writer.TempAppendLine("\t\tData: %v", svcEntityMessage.Data)
	return svcEntityMessage
}

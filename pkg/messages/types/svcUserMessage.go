package messages

import "github.com/pektezol/bitreader"

type SvcUserMessage struct {
	MsgType int8
	Length  int16
	Data    []byte
}

func ParseSvcUserMessage(reader *bitreader.Reader) SvcUserMessage {
	svcUserMessage := SvcUserMessage{
		MsgType: int8(reader.TryReadBits(8)),
		Length:  int16(reader.TryReadBits(12)),
	}
	svcUserMessage.Data = reader.TryReadBitsToSlice(uint64(svcUserMessage.Length))
	return svcUserMessage
}

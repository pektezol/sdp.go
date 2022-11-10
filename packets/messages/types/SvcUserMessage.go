package types

import (
	"github.com/pektezol/bitreader"
)

type SvcUserMessage struct {
	MsgType uint8
	Data    []byte
}

func ParseSvcUserMessage(reader *bitreader.ReaderType) SvcUserMessage {
	msgtype := reader.TryReadInt8()
	length := reader.TryReadBits(12)
	reader.SkipBits(int(length)) // TODO: Read data properly
	return SvcUserMessage{
		MsgType: msgtype,
		//Data:    reader.TryReadBytesToSlice(int(length / 8)),
	}
}

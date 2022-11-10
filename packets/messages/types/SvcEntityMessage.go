package types

import "github.com/pektezol/bitreader"

type SvcEntityMessage struct {
	EntityIndex int16
	ClassId     int16
	Data        []byte
}

func ParseSvcEntityMessage(reader *bitreader.ReaderType) SvcEntityMessage {
	entityindex := reader.TryReadBits(11)
	classid := reader.TryReadBits(9)
	length := reader.TryReadBits(11)
	return SvcEntityMessage{
		EntityIndex: int16(entityindex),
		ClassId:     int16(classid),
		Data:        reader.TryReadBytesToSlice(int(length / 8)),
	}
}

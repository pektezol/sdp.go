package types

import "github.com/pektezol/bitreader"

type SvcGameEventList struct {
	Events int16
	Data   []byte
}

func ParseSvcGameEventList(reader *bitreader.ReaderType) SvcGameEventList {
	events := reader.TryReadBits(9)
	length := reader.TryReadBits(20)
	return SvcGameEventList{
		Events: int16(events),
		Data:   reader.TryReadBytesToSlice(int(length)),
	}
}

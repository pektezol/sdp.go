package types

import "github.com/pektezol/bitreader"

type SvcGameEventList struct {
	Events int16
	Data   []byte
}

func ParseSvcGameEventList(reader *bitreader.ReaderType) SvcGameEventList {
	events := reader.TryReadBits(9)
	length := reader.TryReadBits(20)
	reader.SkipBits(int(length)) // TODO: Read data properly
	return SvcGameEventList{
		Events: int16(events),
		//Data:   reader.TryReadBytesToSlice(int(length)),
	}
}

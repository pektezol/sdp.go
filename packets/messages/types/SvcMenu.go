package types

import "github.com/pektezol/bitreader"

type SvcMenu struct {
	MenuType int16
	Data     []byte
}

func ParseSvcMenu(reader *bitreader.ReaderType) SvcMenu {
	menutype := reader.TryReadInt16()
	length := reader.TryReadInt32()
	return SvcMenu{
		MenuType: int16(menutype),
		Data:     reader.TryReadBytesToSlice(int(length / 8)),
	}
}

package messages

import "github.com/pektezol/bitreader"

type SvcMenu struct {
	MenuType int16
	Length   int32
	Data     []byte
}

func ParseSvcMenu(reader *bitreader.ReaderType) SvcMenu {
	svcMenu := SvcMenu{
		MenuType: int16(reader.TryReadBits(16)),
		Length:   int32(reader.TryReadBits(32)),
	}
	svcMenu.Data = reader.TryReadBitsToSlice(int(svcMenu.Length))
	return svcMenu
}

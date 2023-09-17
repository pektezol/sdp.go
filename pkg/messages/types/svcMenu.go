package messages

import "github.com/pektezol/bitreader"

type SvcMenu struct {
	Type   uint16
	Length uint32
	Data   []byte
}

func ParseSvcMenu(reader *bitreader.Reader) SvcMenu {
	svcMenu := SvcMenu{
		Type:   reader.TryReadUInt16(),
		Length: reader.TryReadUInt32(),
	}
	svcMenu.Data = reader.TryReadBitsToSlice(uint64(svcMenu.Length))
	return svcMenu
}

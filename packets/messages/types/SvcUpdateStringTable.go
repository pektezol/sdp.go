package types

import "github.com/pektezol/bitreader"

type SvcUpdateStringTable struct {
	TableId           int8
	NumChangedEntries int16
	Data              []byte
}

func ParseSvcUpdateStringTable(reader *bitreader.ReaderType) SvcUpdateStringTable {
	svcupdatestringtable := SvcUpdateStringTable{
		TableId: int8(reader.TryReadBits(5)),
	}
	if reader.TryReadBool() {
		svcupdatestringtable.NumChangedEntries = int16(reader.TryReadInt16())
	}
	length := reader.TryReadBits(20)
	reader.SkipBits(int(length)) // TODO: Read data properly
	//svcupdatestringtable.Data = reader.TryReadBytesToSlice(int(length / 8))
	return svcupdatestringtable
}

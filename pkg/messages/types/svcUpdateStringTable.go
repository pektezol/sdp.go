package messages

import "github.com/pektezol/bitreader"

type SvcUpdateStringTable struct {
	TableId           int8
	NumChangedEntries int16
	Length            int32
	Data              []byte
}

func ParseSvcUpdateStringTable(reader *bitreader.ReaderType) SvcUpdateStringTable {
	svcUpdateStringTable := SvcUpdateStringTable{
		TableId: int8(reader.TryReadBits(5)),
	}
	if reader.TryReadBool() {
		svcUpdateStringTable.NumChangedEntries = int16(reader.TryReadBits(16))
	}
	svcUpdateStringTable.Length = int32(reader.TryReadBits(20))
	svcUpdateStringTable.Data = reader.TryReadBitsToSlice(int(svcUpdateStringTable.Length))
	return svcUpdateStringTable
}

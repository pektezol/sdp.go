package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcUpdateStringTable struct {
	TableId           uint8
	NumChangedEntries uint16
	Length            int32
	Data              []byte
}

func ParseSvcUpdateStringTable(reader *bitreader.Reader) SvcUpdateStringTable {
	svcUpdateStringTable := SvcUpdateStringTable{
		TableId: uint8(reader.TryReadBits(5)),
	}
	if reader.TryReadBool() {
		svcUpdateStringTable.NumChangedEntries = reader.TryReadUInt16()
	}
	svcUpdateStringTable.Length = int32(reader.TryReadBits(20))
	svcUpdateStringTable.Data = reader.TryReadBitsToSlice(uint64(svcUpdateStringTable.Length))

	return svcUpdateStringTable
}

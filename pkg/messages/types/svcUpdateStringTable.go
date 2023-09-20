package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	writer.TempAppendLine("\t\tTable ID: %d", svcUpdateStringTable.TableId)
	writer.TempAppendLine("\t\tNumber Of Changed Entries: %d", svcUpdateStringTable.NumChangedEntries)
	return svcUpdateStringTable
}

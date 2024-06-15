package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcUpdateStringTable struct {
	TableId           uint8  `json:"table_id"`
	NumChangedEntries uint16 `json:"num_changed_entries"`
	Length            int32  `json:"length"`
	Data              []byte `json:"data"`
}

func ParseSvcUpdateStringTable(reader *bitreader.Reader, demo *types.Demo) SvcUpdateStringTable {
	svcUpdateStringTable := SvcUpdateStringTable{
		TableId: uint8(reader.TryReadBits(5)),
	}
	if reader.TryReadBool() {
		svcUpdateStringTable.NumChangedEntries = reader.TryReadUInt16()
	}
	svcUpdateStringTable.Length = int32(reader.TryReadBits(20))
	svcUpdateStringTable.Data = reader.TryReadBitsToSlice(uint64(svcUpdateStringTable.Length))
	demo.Writer.TempAppendLine("\t\tTable ID: %d", svcUpdateStringTable.TableId)
	demo.Writer.TempAppendLine("\t\tNumber Of Changed Entries: %d", svcUpdateStringTable.NumChangedEntries)
	return svcUpdateStringTable
}

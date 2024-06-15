package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcPacketEntities struct {
	MaxEntries      uint16 `json:"max_entries"`
	IsDelta         bool   `json:"is_delta"`
	DeltaFrom       int32  `json:"delta_from"`
	BaseLine        bool   `json:"base_line"`
	UpdatedEntries  uint16 `json:"updated_entries"`
	Length          uint32 `json:"length"`
	UpdatedBaseline bool   `json:"updated_baseline"`
	Data            []byte `json:"data"`
}

func ParseSvcPacketEntities(reader *bitreader.Reader, demo *types.Demo) SvcPacketEntities {
	svcPacketEntities := SvcPacketEntities{
		MaxEntries: uint16(reader.TryReadBits(11)),
		IsDelta:    reader.TryReadBool(),
	}
	if svcPacketEntities.IsDelta {
		svcPacketEntities.DeltaFrom = reader.TryReadSInt32()
	} else {
		svcPacketEntities.DeltaFrom = -1
	}
	svcPacketEntities.BaseLine = reader.TryReadBool()
	svcPacketEntities.UpdatedEntries = uint16(reader.TryReadBits(11))
	svcPacketEntities.Length = uint32(reader.TryReadBits(20))
	svcPacketEntities.UpdatedBaseline = reader.TryReadBool()
	svcPacketEntities.Data = reader.TryReadBitsToSlice(uint64(svcPacketEntities.Length))
	demo.Writer.TempAppendLine("\t\tMax Entries: %d", svcPacketEntities.MaxEntries)
	demo.Writer.TempAppendLine("\t\tIs Delta: %t", svcPacketEntities.IsDelta)
	demo.Writer.TempAppendLine("\t\tDelta From: %d", svcPacketEntities.DeltaFrom)
	demo.Writer.TempAppendLine("\t\tBaseline: %t", svcPacketEntities.BaseLine)
	demo.Writer.TempAppendLine("\t\tUpdated Baseline: %t", svcPacketEntities.UpdatedBaseline)
	demo.Writer.TempAppendLine("\t\t%d Updated Entries:", svcPacketEntities.UpdatedEntries)
	demo.Writer.TempAppendLine("\t\tData: %v", svcPacketEntities.Data)
	return svcPacketEntities
}

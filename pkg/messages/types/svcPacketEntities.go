package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcPacketEntities struct {
	MaxEntries      uint16
	IsDelta         bool
	DeltaFrom       int32
	BaseLine        bool
	UpdatedEntries  uint16
	Length          uint32
	UpdatedBaseline bool
	Data            []byte
}

func ParseSvcPacketEntities(reader *bitreader.Reader) SvcPacketEntities {
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
	writer.TempAppendLine("\t\tMax Entries: %d", svcPacketEntities.MaxEntries)
	writer.TempAppendLine("\t\tIs Delta: %t", svcPacketEntities.IsDelta)
	writer.TempAppendLine("\t\tDelta From: %d", svcPacketEntities.DeltaFrom)
	writer.TempAppendLine("\t\tBaseline: %t", svcPacketEntities.BaseLine)
	writer.TempAppendLine("\t\tUpdated Baseline: %t", svcPacketEntities.UpdatedBaseline)
	writer.TempAppendLine("\t\t%d Updated Entries:", svcPacketEntities.UpdatedEntries)
	writer.TempAppendLine("\t\tData: %v", svcPacketEntities.Data)
	return svcPacketEntities
}

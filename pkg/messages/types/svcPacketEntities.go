package messages

import (
	"github.com/pektezol/bitreader"
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
	return svcPacketEntities
}

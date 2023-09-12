package messages

import "github.com/pektezol/bitreader"

type SvcPacketEntities struct {
	MaxEntries      int16
	IsDelta         bool
	DeltaFrom       int32
	BaseLine        bool
	UpdatedEntries  int16
	Length          int32
	UpdatedBaseline bool
	Data            []byte
}

func ParseSvcPacketEntities(reader *bitreader.ReaderType) SvcPacketEntities {
	svcPacketEntities := SvcPacketEntities{
		MaxEntries: int16(reader.TryReadBits(11)),
		IsDelta:    reader.TryReadBool(),
	}
	if svcPacketEntities.IsDelta {
		svcPacketEntities.DeltaFrom = int32(reader.TryReadBits(32))
	}
	svcPacketEntities.BaseLine = reader.TryReadBool()
	svcPacketEntities.UpdatedEntries = int16(reader.TryReadBits(11))
	svcPacketEntities.Length = int32(reader.TryReadBits(20))
	svcPacketEntities.UpdatedBaseline = reader.TryReadBool()
	svcPacketEntities.Data = reader.TryReadBitsToSlice(int(svcPacketEntities.Length))
	return svcPacketEntities
}
package types

import (
	"github.com/pektezol/bitreader"
)

type SvcPacketEntities struct {
	MaxEntries     uint16
	IsDelta        bool
	DeltaFrom      int32
	BaseLine       bool
	UpdatedEntries uint16
	UpdateBaseline bool
	Data           []byte
}

func ParseSvcPacketEntities(reader *bitreader.ReaderType) SvcPacketEntities {
	maxentries := reader.TryReadBits(11)
	isdelta := reader.TryReadBool()
	var deltafrom int32
	if isdelta {
		deltafrom = int32(reader.TryReadInt32())
	}
	baseline := reader.TryReadBool()
	updatedentries := reader.TryReadBits(11)
	length := reader.TryReadBits(20)
	updatebaseline := reader.TryReadBool()
	reader.SkipBits(int(length)) // TODO: Read data properly
	return SvcPacketEntities{
		MaxEntries:     uint16(maxentries),
		IsDelta:        isdelta,
		DeltaFrom:      deltafrom,
		BaseLine:       baseline,
		UpdatedEntries: uint16(updatedentries),
		UpdateBaseline: updatebaseline,
		// Data:           reader.TryReadBytesToSlice(int(length / 8)),
	}
}

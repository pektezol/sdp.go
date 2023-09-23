package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcTempEntities struct {
	NumEntries uint8
	Length     uint32
	Data       []byte
}

func ParseSvcTempEntities(reader *bitreader.Reader) SvcTempEntities {
	svcTempEntities := SvcTempEntities{
		NumEntries: reader.TryReadUInt8(),
		Length:     uint32(reader.TryReadBits(17)),
	}
	svcTempEntities.Data = reader.TryReadBitsToSlice(uint64(svcTempEntities.Length))

	return svcTempEntities
}

package messages

import "github.com/pektezol/bitreader"

type SvcTempEntities struct {
	NumEntries int8
	Length     int32
	Data       []byte
}

func ParseSvcTempEntities(reader *bitreader.Reader) SvcTempEntities {
	svcTempEntities := SvcTempEntities{
		NumEntries: int8(reader.TryReadBits(8)),
		Length:     int32(reader.TryReadBits(17)),
	}
	svcTempEntities.Data = reader.TryReadBitsToSlice(uint64(svcTempEntities.Length))
	return svcTempEntities
}

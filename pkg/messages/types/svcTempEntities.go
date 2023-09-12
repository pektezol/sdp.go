package messages

import "github.com/pektezol/bitreader"

type SvcTempEntities struct {
	NumEntries int8
	Length     int32
	Data       []byte
}

func ParseSvcTempEntities(reader *bitreader.ReaderType) SvcTempEntities {
	svcTempEntities := SvcTempEntities{
		NumEntries: int8(reader.TryReadBits(8)),
		Length:     int32(reader.TryReadBits(17)),
	}
	svcTempEntities.Data = reader.TryReadBitsToSlice(int(svcTempEntities.Length))
	return svcTempEntities
}

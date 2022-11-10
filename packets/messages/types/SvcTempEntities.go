package types

import "github.com/pektezol/bitreader"

type SvcTempEntities struct {
	NumEntries uint8
	Data       []byte
}

func ParseSvcTempEntities(reader *bitreader.ReaderType) SvcTempEntities {
	numentries := reader.TryReadInt8()
	length := reader.TryReadBits(17)
	return SvcTempEntities{
		NumEntries: numentries,
		Data:       reader.TryReadBytesToSlice(int(length / 8)),
	}
}

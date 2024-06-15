package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcTempEntities struct {
	NumEntries uint8  `json:"num_entries"`
	Length     uint32 `json:"length"`
	Data       []byte `json:"data"`
}

func ParseSvcTempEntities(reader *bitreader.Reader, demo *types.Demo) SvcTempEntities {
	svcTempEntities := SvcTempEntities{
		NumEntries: reader.TryReadUInt8(),
		Length:     uint32(reader.TryReadBits(17)),
	}
	svcTempEntities.Data = reader.TryReadBitsToSlice(uint64(svcTempEntities.Length))
	demo.Writer.TempAppendLine("\t\tNumber Of Entries: %d", svcTempEntities.NumEntries)
	demo.Writer.TempAppendLine("\t\tData: %v", svcTempEntities.Data)
	return svcTempEntities
}

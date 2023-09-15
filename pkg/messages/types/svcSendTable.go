package messages

import "github.com/pektezol/bitreader"

type SvcSendTable struct {
	NeedsDecoder int8
	Length       int8
	Props        int32
}

func ParseSvcSendTable(reader *bitreader.Reader) SvcSendTable {
	return SvcSendTable{
		NeedsDecoder: int8(reader.TryReadBits(8)),
		Length:       int8(reader.TryReadBits(8)),
		Props:        int32(reader.TryReadBits(32)),
	}
}

package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcSendTable struct {
	NeedsDecoder bool
	Length       uint8
	Props        uint32
}

func ParseSvcSendTable(reader *bitreader.Reader) SvcSendTable {
	svcSendTable := SvcSendTable{
		NeedsDecoder: reader.TryReadBool(),
		Length:       reader.TryReadUInt8(),
	}
	svcSendTable.Props = uint32(reader.TryReadBits(uint64(svcSendTable.Length)))

	return svcSendTable
}

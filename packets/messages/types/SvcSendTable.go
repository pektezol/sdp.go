package types

import "github.com/pektezol/bitreader"

type SvcSendTable struct {
	NeedsDecoder bool
	Length       uint8
	Props        int32
}

func ParseSvcSendTable(reader *bitreader.ReaderType) SvcSendTable {
	return SvcSendTable{
		NeedsDecoder: reader.TryReadBool(),
		Length:       reader.TryReadInt8(),
	}
	// No one cares about SvcSendTable
}

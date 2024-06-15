package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcSendTable struct {
	NeedsDecoder bool   `json:"needs_decoder"`
	Length       uint8  `json:"length"`
	Props        uint32 `json:"props"`
}

func ParseSvcSendTable(reader *bitreader.Reader, demo *types.Demo) SvcSendTable {
	svcSendTable := SvcSendTable{
		NeedsDecoder: reader.TryReadBool(),
		Length:       reader.TryReadUInt8(),
	}
	svcSendTable.Props = uint32(reader.TryReadBits(uint64(svcSendTable.Length)))
	demo.Writer.TempAppendLine("\t\tNeeds Decoder: %t", svcSendTable.NeedsDecoder)
	demo.Writer.TempAppendLine("\t\tLength: %d", svcSendTable.Length)
	demo.Writer.TempAppendLine("\t\tProps: %d", svcSendTable.Props)
	return svcSendTable
}

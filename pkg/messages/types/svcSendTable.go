package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
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
	writer.TempAppendLine("\t\tNeeds Decoder: %t", svcSendTable.NeedsDecoder)
	writer.TempAppendLine("\t\tLength: %d", svcSendTable.Length)
	writer.TempAppendLine("\t\tProps: %d", svcSendTable.Props)
	return svcSendTable
}

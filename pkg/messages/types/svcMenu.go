package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcMenu struct {
	Type   uint16 `json:"type"`
	Length uint32 `json:"length"`
	Data   []byte `json:"data"`
}

func ParseSvcMenu(reader *bitreader.Reader, demo *types.Demo) SvcMenu {
	svcMenu := SvcMenu{
		Type:   reader.TryReadUInt16(),
		Length: reader.TryReadUInt32(),
	}
	svcMenu.Data = reader.TryReadBitsToSlice(uint64(svcMenu.Length))
	demo.Writer.TempAppendLine("\t\tType: %d", svcMenu.Type)
	demo.Writer.TempAppendLine("\t\tData: %v", svcMenu.Data)
	return svcMenu
}

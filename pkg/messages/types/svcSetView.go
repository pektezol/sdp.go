package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcSetView struct {
	EntityIndex uint16 `json:"entity_index"`
}

func ParseSvcSetView(reader *bitreader.Reader, demo *types.Demo) SvcSetView {
	svcSetView := SvcSetView{
		EntityIndex: uint16(reader.TryReadBits(11)),
	}
	demo.Writer.TempAppendLine("\t\tEntity Index: %d", svcSetView.EntityIndex)
	return svcSetView
}

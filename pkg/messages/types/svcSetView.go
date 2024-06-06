package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type SvcSetView struct {
	EntityIndex uint16
}

func ParseSvcSetView(reader *bitreader.Reader) SvcSetView {
	svcSetView := SvcSetView{
		EntityIndex: uint16(reader.TryReadBits(11)),
	}
	writer.TempAppendLine("\t\tEntity Index: %d", svcSetView.EntityIndex)
	return svcSetView
}

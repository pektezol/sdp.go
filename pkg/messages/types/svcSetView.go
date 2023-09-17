package messages

import "github.com/pektezol/bitreader"

type SvcSetView struct {
	EntityIndex uint16
}

func ParseSvcSetView(reader *bitreader.Reader) SvcSetView {
	return SvcSetView{
		EntityIndex: uint16(reader.TryReadBits(11)),
	}
}

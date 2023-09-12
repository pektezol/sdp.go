package messages

import "github.com/pektezol/bitreader"

type SvcSetView struct {
	EntityIndex int16
}

func ParseSvcSetView(reader *bitreader.ReaderType) SvcSetView {
	return SvcSetView{
		EntityIndex: int16(reader.TryReadBits(11)),
	}
}

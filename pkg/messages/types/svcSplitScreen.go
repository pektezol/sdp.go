package messages

import "github.com/pektezol/bitreader"

type SvcSplitScreen struct {
	RemoveUser bool
	Length     uint16
	Data       []byte
}

func ParseSvcSplitScreen(reader *bitreader.Reader) SvcSplitScreen {
	svcSplitScreen := SvcSplitScreen{
		RemoveUser: reader.TryReadBool(),
		Length:     uint16(reader.TryReadBits(11)),
	}
	svcSplitScreen.Data = reader.TryReadBitsToSlice(uint64(svcSplitScreen.Length))
	return svcSplitScreen
}

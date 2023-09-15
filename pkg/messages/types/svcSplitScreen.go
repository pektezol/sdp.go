package messages

import "github.com/pektezol/bitreader"

type SvcSplitScreen struct {
	Unk    bool
	Length int16
	Data   []byte
}

func ParseSvcSplitScreen(reader *bitreader.Reader) SvcSplitScreen {
	svcSplitScreen := SvcSplitScreen{
		Unk:    reader.TryReadBool(),
		Length: int16(reader.TryReadBits(11)),
	}
	svcSplitScreen.Data = reader.TryReadBitsToSlice(uint64(svcSplitScreen.Length))
	return svcSplitScreen
}

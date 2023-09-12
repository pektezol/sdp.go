package messages

import "github.com/pektezol/bitreader"

type SvcSplitScreen struct {
	Unk    bool
	Length int16
	Data   []byte
}

func ParseSvcSplitScreen(reader *bitreader.ReaderType) SvcSplitScreen {
	svcSplitScreen := SvcSplitScreen{
		Unk:    reader.TryReadBool(),
		Length: int16(reader.TryReadBits(11)),
	}
	svcSplitScreen.Data = reader.TryReadBitsToSlice(int(svcSplitScreen.Length))
	return svcSplitScreen
}

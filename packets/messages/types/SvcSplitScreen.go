package types

import "github.com/pektezol/bitreader"

type SvcSplitScreen struct {
	Unk  bool
	Data []byte
}

func ParseSvcSplitScreen(reader *bitreader.ReaderType) SvcSplitScreen {
	unk := reader.TryReadBool()
	length := reader.TryReadBits(11)
	return SvcSplitScreen{
		Unk:  unk,
		Data: reader.TryReadBytesToSlice(int(length / 8)),
	}
}

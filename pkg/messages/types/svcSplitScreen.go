package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

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
	writer.TempAppendLine("\t\tRemove User: %t", svcSplitScreen.RemoveUser)
	writer.TempAppendLine("\t\tData: %v", svcSplitScreen.Data)
	return svcSplitScreen
}

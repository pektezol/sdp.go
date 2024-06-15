package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcSplitScreen struct {
	RemoveUser bool   `json:"remove_user"`
	Length     uint16 `json:"length"`
	Data       []byte `json:"data"`
}

func ParseSvcSplitScreen(reader *bitreader.Reader, demo *types.Demo) SvcSplitScreen {
	svcSplitScreen := SvcSplitScreen{
		RemoveUser: reader.TryReadBool(),
		Length:     uint16(reader.TryReadBits(11)),
	}
	svcSplitScreen.Data = reader.TryReadBitsToSlice(uint64(svcSplitScreen.Length))
	demo.Writer.TempAppendLine("\t\tRemove User: %t", svcSplitScreen.RemoveUser)
	demo.Writer.TempAppendLine("\t\tData: %v", svcSplitScreen.Data)
	return svcSplitScreen
}

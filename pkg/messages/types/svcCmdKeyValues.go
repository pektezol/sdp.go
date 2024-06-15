package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcCmdKeyValues struct {
	Length uint32 `json:"length"`
	Data   []byte `json:"data"`
}

func ParseSvcCmdKeyValues(reader *bitreader.Reader, demo *types.Demo) SvcCmdKeyValues {
	svcCmdKeyValues := SvcCmdKeyValues{
		Length: reader.TryReadUInt32(),
	}
	svcCmdKeyValues.Data = reader.TryReadBytesToSlice(uint64(svcCmdKeyValues.Length))
	demo.Writer.TempAppendLine("\t\tData: %v", svcCmdKeyValues.Data)
	return svcCmdKeyValues
}

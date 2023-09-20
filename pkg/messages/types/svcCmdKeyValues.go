package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcCmdKeyValues struct {
	Length uint32
	Data   []byte
}

func ParseSvcCmdKeyValues(reader *bitreader.Reader) SvcCmdKeyValues {
	svcCmdKeyValues := SvcCmdKeyValues{
		Length: reader.TryReadUInt32(),
	}
	svcCmdKeyValues.Data = reader.TryReadBytesToSlice(uint64(svcCmdKeyValues.Length))
	writer.TempAppendLine("\t\tData: %v", svcCmdKeyValues.Data)
	return svcCmdKeyValues
}

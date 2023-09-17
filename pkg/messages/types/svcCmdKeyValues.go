package messages

import "github.com/pektezol/bitreader"

type SvcCmdKeyValues struct {
	Length uint32
	Data   []byte
}

func ParseSvcCmdKeyValues(reader *bitreader.Reader) SvcCmdKeyValues {
	svcCmdKeyValues := SvcCmdKeyValues{
		Length: reader.TryReadUInt32(),
	}
	svcCmdKeyValues.Data = reader.TryReadBytesToSlice(uint64(svcCmdKeyValues.Length))
	return svcCmdKeyValues
}

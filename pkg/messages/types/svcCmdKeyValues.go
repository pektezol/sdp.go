package messages

import "github.com/pektezol/bitreader"

type SvcCmdKeyValues struct {
	Length int32
	Data   []byte
}

func ParseSvcCmdKeyValues(reader *bitreader.Reader) SvcCmdKeyValues {
	svcCmdKeyValues := SvcCmdKeyValues{
		Length: int32(reader.TryReadBits(32)),
	}
	svcCmdKeyValues.Data = reader.TryReadBytesToSlice(uint64(svcCmdKeyValues.Length))
	return svcCmdKeyValues
}

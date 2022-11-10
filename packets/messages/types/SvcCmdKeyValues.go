package types

import "github.com/pektezol/bitreader"

type SvcCmdKeyValues struct {
	Data []byte
}

func ParseSvcCmdKeyValues(reader *bitreader.ReaderType) SvcCmdKeyValues {
	length := reader.TryReadInt32()
	return SvcCmdKeyValues{
		Data: reader.TryReadBytesToSlice(int(length)),
	}
}

package types

import "github.com/pektezol/bitreader"

type SvcVoiceData struct {
	Client    int8
	Proximity int8
	Data      []byte
}

func ParseSvcVoiceData(reader *bitreader.ReaderType) SvcVoiceData {
	svcvoicedata := SvcVoiceData{
		Client:    int8(reader.TryReadInt8()),
		Proximity: int8(reader.TryReadInt8()),
	}
	length := reader.TryReadInt16()
	svcvoicedata.Data = reader.TryReadBytesToSlice(int(length / 8))
	return svcvoicedata
}

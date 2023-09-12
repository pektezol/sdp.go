package messages

import "github.com/pektezol/bitreader"

type SvcVoiceData struct {
	Client    int8
	Proximity int8
	Length    int16
	Data      []byte
}

func ParseSvcVoiceData(reader *bitreader.ReaderType) SvcVoiceData {
	svcVoiceData := SvcVoiceData{
		Client:    int8(reader.TryReadBits(8)),
		Proximity: int8(reader.TryReadBits(8)),
		Length:    int16(reader.TryReadBits(16)),
	}
	svcVoiceData.Data = reader.TryReadBitsToSlice(int(svcVoiceData.Length))
	return svcVoiceData
}

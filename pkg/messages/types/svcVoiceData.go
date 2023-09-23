package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcVoiceData struct {
	FromClient uint8
	Proximity  bool
	Length     int16
	Data       []byte
}

func ParseSvcVoiceData(reader *bitreader.Reader) SvcVoiceData {
	svcVoiceData := SvcVoiceData{
		FromClient: reader.TryReadUInt8(),
	}
	proximity := reader.TryReadUInt8()
	if proximity != 0 {
		svcVoiceData.Proximity = true
	}
	svcVoiceData.Data = reader.TryReadBitsToSlice(uint64(svcVoiceData.Length))

	return svcVoiceData
}

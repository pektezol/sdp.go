package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	writer.TempAppendLine("\t\tFrom Client: %d", svcVoiceData.FromClient)
	writer.TempAppendLine("\t\tProximity: %t", svcVoiceData.Proximity)
	writer.TempAppendLine("\t\tData: %v", svcVoiceData.Data)
	return svcVoiceData
}

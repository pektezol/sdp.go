package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcVoiceData struct {
	FromClient uint8  `json:"from_client"`
	Proximity  bool   `json:"proximity"`
	Length     int16  `json:"length"`
	Data       []byte `json:"data"`
}

func ParseSvcVoiceData(reader *bitreader.Reader, demo *types.Demo) SvcVoiceData {
	svcVoiceData := SvcVoiceData{
		FromClient: reader.TryReadUInt8(),
	}
	proximity := reader.TryReadUInt8()
	if proximity != 0 {
		svcVoiceData.Proximity = true
	}
	svcVoiceData.Data = reader.TryReadBitsToSlice(uint64(svcVoiceData.Length))
	demo.Writer.TempAppendLine("\t\tFrom Client: %d", svcVoiceData.FromClient)
	demo.Writer.TempAppendLine("\t\tProximity: %t", svcVoiceData.Proximity)
	demo.Writer.TempAppendLine("\t\tData: %v", svcVoiceData.Data)
	return svcVoiceData
}

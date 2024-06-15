package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcVoiceInit struct {
	Codec      string `json:"codec"`
	Quality    uint8  `json:"quality"`
	SampleRate int32  `json:"sample_rate"`
}

func ParseSvcVoiceInit(reader *bitreader.Reader, demo *types.Demo) SvcVoiceInit {
	svcVoiceInit := SvcVoiceInit{
		Codec:   reader.TryReadString(),
		Quality: reader.TryReadUInt8(),
	}
	if svcVoiceInit.Quality == 0b11111111 {
		svcVoiceInit.SampleRate = reader.TryReadSInt32()
	} else {
		if svcVoiceInit.Codec == "vaudio_celt" {
			svcVoiceInit.SampleRate = 22050
		} else {
			svcVoiceInit.SampleRate = 11025
		}
	}
	demo.Writer.TempAppendLine("\t\tCodec: %s", svcVoiceInit.Codec)
	demo.Writer.TempAppendLine("\t\tQuality: %d", svcVoiceInit.Quality)
	demo.Writer.TempAppendLine("\t\tSample Rate: %d", svcVoiceInit.SampleRate)
	return svcVoiceInit
}

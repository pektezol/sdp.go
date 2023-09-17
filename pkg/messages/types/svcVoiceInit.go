package messages

import "github.com/pektezol/bitreader"

type SvcVoiceInit struct {
	Codec      string
	Quality    uint8
	SampleRate int32
}

func ParseSvcVoiceInit(reader *bitreader.Reader) SvcVoiceInit {
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
	return svcVoiceInit
}

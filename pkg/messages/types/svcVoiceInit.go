package messages

import "github.com/pektezol/bitreader"

type SvcVoiceInit struct {
	Codec   string
	Quality uint8
	Unk     float32
}

func ParseSvcVoiceInit(reader *bitreader.ReaderType) SvcVoiceInit {
	svcVoiceInit := SvcVoiceInit{
		Codec:   reader.TryReadString(),
		Quality: uint8(reader.TryReadBits(8)),
	}
	if svcVoiceInit.Quality == 0b11111111 {
		svcVoiceInit.Unk = reader.TryReadFloat32()
	}
	return svcVoiceInit
}

package types

import "github.com/pektezol/bitreader"

type SvcVoiceInit struct {
	Codec   string
	Quality uint8
	Unk     float32
}

func ParseSvcVoiceInit(reader *bitreader.ReaderType) SvcVoiceInit {
	svcvoiceinit := SvcVoiceInit{
		Codec:   reader.TryReadString(),
		Quality: reader.TryReadInt8(),
	}
	if svcvoiceinit.Quality == 255 {
		svcvoiceinit.Unk = reader.TryReadFloat32()
	}
	return svcvoiceinit
}

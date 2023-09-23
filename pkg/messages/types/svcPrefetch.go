package messages

import (
	"github.com/pektezol/bitreader"
)

type SvcPrefetch struct {
	SoundIndex int16
}

func ParseSvcPrefetch(reader *bitreader.Reader) SvcPrefetch {
	svcPrefetch := SvcPrefetch{
		SoundIndex: int16(reader.TryReadBits(13)),
	}

	return svcPrefetch
}

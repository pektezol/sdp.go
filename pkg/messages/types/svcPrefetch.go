package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcPrefetch struct {
	SoundIndex int16 `json:"sound_index"`
}

func ParseSvcPrefetch(reader *bitreader.Reader, demo *types.Demo) SvcPrefetch {
	svcPrefetch := SvcPrefetch{
		SoundIndex: int16(reader.TryReadBits(13)),
	}
	demo.Writer.TempAppendLine("\t\tSound Index: %d", svcPrefetch.SoundIndex)
	return svcPrefetch
}

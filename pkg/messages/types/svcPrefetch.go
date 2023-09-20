package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type SvcPrefetch struct {
	SoundIndex int16
}

func ParseSvcPrefetch(reader *bitreader.Reader) SvcPrefetch {
	svcPrefetch := SvcPrefetch{
		SoundIndex: int16(reader.TryReadBits(13)),
	}
	writer.TempAppendLine("\t\tSound Index: %d", svcPrefetch.SoundIndex)
	return svcPrefetch
}

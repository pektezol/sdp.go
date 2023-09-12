package messages

import "github.com/pektezol/bitreader"

type SvcPrefetch struct {
	SoundIndex int16
}

func ParseSvcPrefetch(reader *bitreader.ReaderType) SvcPrefetch {
	return SvcPrefetch{
		SoundIndex: int16(reader.TryReadBits(13)),
	}
}

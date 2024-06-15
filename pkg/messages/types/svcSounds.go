package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type SvcSounds struct {
	ReliableSound bool   `json:"reliable_sound"`
	SoundCount    uint8  `json:"sound_count"`
	Length        uint16 `json:"length"`
	Data          []byte `json:"data"`
}

func ParseSvcSounds(reader *bitreader.Reader, demo *types.Demo) SvcSounds {
	svcSounds := SvcSounds{
		ReliableSound: reader.TryReadBool(),
	}
	if svcSounds.ReliableSound {
		svcSounds.SoundCount = 1
		svcSounds.Length = uint16(reader.TryReadUInt8())
	} else {
		svcSounds.SoundCount = reader.TryReadUInt8()
		svcSounds.Length = reader.TryReadUInt16()
	}
	svcSounds.Data = reader.TryReadBitsToSlice(uint64(svcSounds.Length))
	demo.Writer.TempAppendLine("\t\tReliable Sound: %t", svcSounds.ReliableSound)
	demo.Writer.TempAppendLine("\t\tSound Count: %d", svcSounds.SoundCount)
	demo.Writer.TempAppendLine("\t\tData: %v", svcSounds.Data)
	return svcSounds
}

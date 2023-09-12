package messages

import "github.com/pektezol/bitreader"

type SvcSounds struct {
	ReliableSound bool
	Size          int8
	Length        int16
	Data          []byte
}

func ParseSvcSounds(reader *bitreader.ReaderType) SvcSounds {
	svcSounds := SvcSounds{
		ReliableSound: reader.TryReadBool(),
	}
	if svcSounds.ReliableSound {
		svcSounds.Size = 1
		svcSounds.Length = int16(reader.TryReadBits(8))
	} else {
		svcSounds.Size = int8(reader.TryReadBits(8))
		svcSounds.Length = int16(reader.TryReadBits(16))
	}
	svcSounds.Data = reader.TryReadBitsToSlice(int(svcSounds.Length))
	return svcSounds
}

package types

import "github.com/pektezol/bitreader"

type SvcSounds struct {
	ReliableSound bool
	Size          int8
	Data          []byte
}

func ParseSvcSounds(reader *bitreader.ReaderType) SvcSounds {
	reliablesound := reader.TryReadBool()
	var size int8
	var length int16
	if reliablesound {
		size = 1
	} else {
		size = int8(reader.TryReadInt8())
	}
	if reliablesound {
		length = int16(reader.TryReadInt8())
	} else {
		length = int16(reader.TryReadInt16())
	}
	data := reader.TryReadBytesToSlice(int(length / 8))
	return SvcSounds{
		ReliableSound: reliablesound,
		Size:          size,
		Data:          data,
	}
}

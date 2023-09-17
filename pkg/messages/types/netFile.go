package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
)

type NetFile struct {
	TransferId uint32
	FileName   string
	FileFlags  string
}

type NetFileFlags int

const (
	ENetFileFlagsNone          NetFileFlags = 0
	ENetFileFlagsFileRequested NetFileFlags = 1
	ENetFileFlagsUnknown       NetFileFlags = 1 << 1
)

func (netFileFlags NetFileFlags) String() string {
	switch netFileFlags {
	case ENetFileFlagsNone:
		return "None"
	case ENetFileFlagsFileRequested:
		return "FileRequested"
	case ENetFileFlagsUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("%d", int(netFileFlags))
	}
}

func ParseNetFile(reader *bitreader.Reader) NetFile {
	return NetFile{
		TransferId: reader.TryReadUInt32(),
		FileName:   reader.TryReadString(),
		FileFlags:  NetFileFlags(reader.TryReadBits(2)).String(),
	}
}

package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
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
	netFile := NetFile{
		TransferId: reader.TryReadUInt32(),
		FileName:   reader.TryReadString(),
		FileFlags:  NetFileFlags(reader.TryReadBits(2)).String(),
	}
	writer.TempAppendLine("\t\tTransfer ID: %d", netFile.TransferId)
	writer.TempAppendLine("\t\tFile Name: %s", netFile.FileName)
	writer.TempAppendLine("\t\tFile Flags: %s", netFile.FileFlags)
	return netFile
}

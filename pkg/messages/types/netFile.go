package messages

import (
	"fmt"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetFile struct {
	TransferId uint32 `json:"transfer_id"`
	FileName   string `json:"file_name"`
	FileFlags  string `json:"file_flags"`
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

func ParseNetFile(reader *bitreader.Reader, demo *types.Demo) NetFile {
	netFile := NetFile{
		TransferId: reader.TryReadUInt32(),
		FileName:   reader.TryReadString(),
		FileFlags:  NetFileFlags(reader.TryReadBits(2)).String(),
	}
	demo.Writer.TempAppendLine("\t\tTransfer ID: %d", netFile.TransferId)
	demo.Writer.TempAppendLine("\t\tFile Name: %s", netFile.FileName)
	demo.Writer.TempAppendLine("\t\tFile Flags: %s", netFile.FileFlags)
	return netFile
}

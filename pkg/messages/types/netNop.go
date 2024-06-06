package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/writer"
)

type NetNop struct{}

func ParseNetNop(reader *bitreader.Reader) NetNop {
	writer.TempAppendLine("\t\t{}")
	return NetNop{}
}

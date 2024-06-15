package messages

import (
	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetNop struct{}

func ParseNetNop(reader *bitreader.Reader, demo *types.Demo) NetNop {
	demo.Writer.TempAppendLine("\t\t{}")
	return NetNop{}
}

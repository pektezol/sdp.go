package messages

import (
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/sdp.go/pkg/types"
)

type NetStringCmd struct {
	Command string `json:"command"`
}

func ParseNetStringCmd(reader *bitreader.Reader, demo *types.Demo) NetStringCmd {
	netStringCmd := NetStringCmd{
		Command: reader.TryReadString(),
	}
	demo.Writer.TempAppendLine("\t\tCommand: \"%s\"", strings.TrimSpace(netStringCmd.Command))
	return netStringCmd
}

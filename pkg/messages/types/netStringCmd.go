package messages

import (
	"strings"

	"github.com/pektezol/bitreader"
	"github.com/pektezol/demoparser/pkg/writer"
)

type NetStringCmd struct {
	Command string
}

func ParseNetStringCmd(reader *bitreader.Reader) NetStringCmd {
	netStringCmd := NetStringCmd{
		Command: reader.TryReadString(),
	}
	writer.TempAppendLine("\t\tCommand: \"%s\"", strings.TrimSpace(netStringCmd.Command))
	return netStringCmd
}

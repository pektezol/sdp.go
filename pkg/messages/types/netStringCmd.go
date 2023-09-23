package messages

import (
	"github.com/pektezol/bitreader"
)

type NetStringCmd struct {
	Command string
}

func ParseNetStringCmd(reader *bitreader.Reader) NetStringCmd {
	netStringCmd := NetStringCmd{
		Command: reader.TryReadString(),
	}

	return netStringCmd
}

package messages

import "github.com/pektezol/bitreader"

type NetStringCmd struct {
	Command string
}

func ParseNetStringCmd(reader *bitreader.Reader) NetStringCmd {
	return NetStringCmd{
		Command: reader.TryReadString(),
	}
}

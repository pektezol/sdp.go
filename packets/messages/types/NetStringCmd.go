package types

import "github.com/pektezol/bitreader"

type NetStringCmd struct {
	Command string
}

func ParseNetStringCmd(reader *bitreader.ReaderType) NetStringCmd {
	return NetStringCmd{
		Command: reader.TryReadString(),
	}
}
